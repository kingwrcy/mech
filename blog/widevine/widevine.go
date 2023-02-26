package widevine

import (
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/pem"
   "errors"
   "github.com/aead/cmac"
   "google.golang.org/protobuf/proto"
   "time"
)

type CDM struct {
   
   clientID   []byte
   privacyMode             bool
   privateKey *rsa.PrivateKey
   sessionID  [32]byte
   signedDeviceCertificate SignedDeviceCertificate
   widevineCencHeader      WidevineCencHeader
}

type Key struct {
   ID    []byte
   Type  License_KeyContainer_KeyType
   Value []byte
}

// Creates a new CDM object with the specified device information.
func NewCDM(privateKey, clientID, initData []byte) (CDM, error) {
   block, _ := pem.Decode(privateKey)
   if block == nil || (block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY") {
      return CDM{}, errors.New("failed to decode device private key")
   }

   keyParsed, err := x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      // if PCKS1 doesn't work, try PCKS8
      pcks8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
      if err != nil {
         return CDM{}, err
      }
      keyParsed = pcks8Key.(*rsa.PrivateKey)
   }

   var widevineCencHeader WidevineCencHeader
   if len(initData) < 32 {
      return CDM{}, errors.New("initData not long enough")
   }
   if err := proto.Unmarshal(initData[32:], &widevineCencHeader); err != nil {
      return CDM{}, err
   }
   return CDM{
      clientID:   clientID,
      privateKey: keyParsed,
      widevineCencHeader: widevineCencHeader,
   }, nil
}

// Generates the license request data.  This is sent to the license server via
// HTTP POST and the server in turn returns the license response.
func (c *CDM) GetLicenseRequest() ([]byte, error) {
   var licenseRequest SignedLicenseRequest
   licenseRequest.Msg = new(LicenseRequest)
   licenseRequest.Msg.ContentId = new(LicenseRequest_ContentIdentification)
   licenseRequest.Msg.ContentId.CencId = new(LicenseRequest_ContentIdentification_CENC)

   // this is probably really bad for the GC but protobuf uses pointers for optional
   // fields so it is necessary and this is not a long running program
   {
      v := SignedLicenseRequest_LICENSE_REQUEST
      licenseRequest.Type = &v
   }

   licenseRequest.Msg.ContentId.CencId.Pssh = &c.widevineCencHeader

   {
      v := LicenseType_DEFAULT
      licenseRequest.Msg.ContentId.CencId.LicenseType = &v
   }

   licenseRequest.Msg.ContentId.CencId.RequestId = c.sessionID[:]

   {
      v := LicenseRequest_NEW
      licenseRequest.Msg.Type = &v
   }

   {
      v := uint32(time.Now().Unix())
      licenseRequest.Msg.RequestTime = &v
   }

   {
      v := ProtocolVersion_CURRENT
      licenseRequest.Msg.ProtocolVersion = &v
   }
   licenseRequest.Msg.ClientId = new(ClientIdentification)
   err := proto.Unmarshal(c.clientID, licenseRequest.Msg.ClientId)
   if err != nil {
      return nil, err
   }
   {
      data, err := proto.Marshal(licenseRequest.Msg)
      if err != nil {
         return nil, err
      }
      hash := sha1.Sum(data)
      if licenseRequest.Signature, err = rsa.SignPSS(
         no_operation{},
         c.privateKey, crypto.SHA1, hash[:],
         &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
      ); err != nil {
         return nil, err
      }
   }

   return proto.Marshal(&licenseRequest)
}

type no_operation struct{}

func (no_operation) Read(buf []byte) (int, error) {
   return len(buf), nil
}

// Retrieves the keys from the license response data.  These keys can be
// used to decrypt the DASH-MP4.
func (c *CDM) GetLicenseKeys(licenseRequest []byte, licenseResponse []byte) (keys []Key, err error) {
   var license SignedLicense
   if err = proto.Unmarshal(licenseResponse, &license); err != nil {
      return
   }

   var licenseRequestParsed SignedLicenseRequest
   if err = proto.Unmarshal(licenseRequest, &licenseRequestParsed); err != nil {
      return
   }
   licenseRequestMsg, err := proto.Marshal(licenseRequestParsed.Msg)
   if err != nil {
      return
   }
   sessionKey, err := rsa.DecryptOAEP(
      sha1.New(),
      nil,
      c.privateKey, license.SessionKey, nil,
   )
   if err != nil {
      return
   }

   sessionKeyBlock, err := aes.NewCipher(sessionKey)
   if err != nil {
      return
   }

   encryptionKey := []byte{1, 'E', 'N', 'C', 'R', 'Y', 'P', 'T', 'I', 'O', 'N', 0}
   encryptionKey = append(encryptionKey, licenseRequestMsg...)
   encryptionKey = append(encryptionKey, []byte{0, 0, 0, 0x80}...)
   encryptionKeyCmac, err := cmac.Sum(encryptionKey, sessionKeyBlock, sessionKeyBlock.BlockSize())
   if err != nil {
      return
   }
   encryptionKeyCipher, err := aes.NewCipher(encryptionKeyCmac)
   if err != nil {
      return
   }

   unpad := func(b []byte) []byte {
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
   for _, key := range license.Msg.Key {
      decrypter := cipher.NewCBCDecrypter(encryptionKeyCipher, key.Iv)
      decryptedKey := make([]byte, len(key.Key))
      decrypter.CryptBlocks(decryptedKey, key.Key)
      keys = append(keys, Key{
         ID:    key.Id,
         Type:  *key.Type,
         Value: unpad(decryptedKey),
      })
   }

   return
}
