package widevine

import (
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/pem"
   "github.com/chmike/cmac-go"
   "google.golang.org/protobuf/proto"
)

type no_operation struct{}

func (no_operation) Read(buf []byte) (int, error) {
   return len(buf), nil
}

func unpad(b []byte) []byte {
   if len(b) == 0 {
      return b
   }
   // pks padding is designed so that the value of all the padding bytes is
   // the number of padding bytes repeated so to figure out how many
   // padding bytes there are we can just look at the value of the last
   // byte
   // i.e if there are 6 padding bytes then it will look at like
   // <data> 0x6 0x6 0x6 0x6 0x6 0x6
   count := int(b[len(b)-1])
   return b[0 : len(b)-count]
}

type Container struct {
   Type  License_KeyContainer_KeyType
   Value []byte
}

type Module struct {
   client_ID   []byte
   privateKey *rsa.PrivateKey
   widevineCencHeader      WidevineCencHeader
}

// Creates a new Module object with the specified device information.
func New_Module(privateKey, client_ID, init_data []byte) (*Module, error) {
   block, _ := pem.Decode(privateKey)
   var (
      err error
      mod Module
   )
   mod.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   if err := proto.Unmarshal(init_data[32:], &mod.widevineCencHeader); err != nil {
      return nil, err
   }
   mod.client_ID = client_ID
   return &mod, nil
}

// Generates the license request data.  This is sent to the license server via
// HTTP POST and the server in turn returns the license response.
func (m *Module) signed_request() ([]byte, error) {
   var license_req SignedLicenseRequest
   license_req.Msg = new(LicenseRequest)
   license_req.Msg.ClientId = new(ClientIdentification)
   license_req.Msg.ContentId = new(LicenseRequest_ContentIdentification)
   license_req.Msg.ContentId.CencId = new(LicenseRequest_ContentIdentification_CENC)
   license_req.Msg.ContentId.CencId.Pssh = &m.widevineCencHeader
   err := proto.Unmarshal(m.client_ID, license_req.Msg.ClientId)
   if err != nil {
      return nil, err
   }
   data, err := proto.Marshal(license_req.Msg)
   if err != nil {
      return nil, err
   }
   hash := sha1.Sum(data)
   license_req.Signature, err = rsa.SignPSS(
      no_operation{},
      m.privateKey,
      crypto.SHA1,
      hash[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   return proto.Marshal(&license_req)
}

// Retrieves the keys from the license response data. These keys can be used to
// decrypt the DASH-MP4.
func (m *Module) GetLicenseKeys(license_req []byte, license_res []byte) ([]Container, error) {
   var signed_response SignedLicense
   err := proto.Unmarshal(license_res, &signed_response)
   if err != nil {
      return nil, err
   }
   session_key, err := rsa.DecryptOAEP(
      sha1.New(), nil, m.privateKey, signed_response.SessionKey, nil,
   )
   if err != nil {
      return nil, err
   }
   var request_parsed SignedLicenseRequest
   if err = proto.Unmarshal(license_req, &request_parsed); err != nil {
      return nil, err
   }
   request_msg, err := proto.Marshal(request_parsed.Msg)
   if err != nil {
      return nil, err
   }
   enc_key := []byte{1, 'E', 'N', 'C', 'R', 'Y', 'P', 'T', 'I', 'O', 'N', 0}
   enc_key = append(enc_key, request_msg...)
   enc_key = append(enc_key, []byte{0, 0, 0, 0x80}...)
   // CMAC
   key_CMAC, err := cmac.New(aes.NewCipher, session_key)
   if err != nil {
      return nil, err
   }
   key_CMAC.Write(enc_key)
   key_cipher, err := aes.NewCipher(key_CMAC.Sum(nil))
   if err != nil {
      return nil, err
   }
   var keys []Container
   for _, key := range signed_response.Msg.Key {
      cipher.NewCBCDecrypter(key_cipher, key.Iv).CryptBlocks(key.Key, key.Key)
      keys = append(keys, Container{
         Type:  *key.Type,
         Value: unpad(key.Key),
      })
   }
   return keys, nil
}
