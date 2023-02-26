package widevine

import (
   "2a.pages.dev/mech/roku"
   "bytes"
   "encoding/base64"
   "encoding/hex"
   "io"
   "net/http"
   "os"
   "testing"
   "time"
)

const (
   episode = iota
   movie
)

type key struct {
   media_type int
   status int
}

var tests = map[key]struct {
   key string
   playback_ID string
   init_data string
} {
   // therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
   {episode, 500}: {
      init_data: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg==",
      key: "e258b67d75420066c8424bd142f84565",
      playback_ID: "105c41ea75775968b670fbb26978ed76",
   },
   // therokuchannel.roku.com/watch/12cdc6139bb858f4a597a1aa51a1e37d
   {episode, 200}: {
      init_data: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQmKI1QsJl+5Fx3Kpnu/oGTBoKaW50ZXJ0cnVzdCIBKg==",
      key: "17e71e939c2a541142112c376cf73cc4",
      playback_ID: "12cdc6139bb858f4a597a1aa51a1e37d",
   },
}

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      init_data, err := base64.StdEncoding.DecodeString(test.init_data)
      if err != nil {
         t.Fatal(err)
      }
      mod, err := NewCDM(private_key, client_ID, init_data)
      if err != nil {
         t.Fatal(err)
      }
      req_body, err := mod.GetLicenseRequest()
      if err != nil {
         t.Fatal(err)
      }
      site, err := roku.New_Cross_Site()
      if err != nil {
         t.Fatal(err)
      }
      play, err := site.Playback(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      req, err := http.NewRequest(
         "POST", play.DRM.Widevine.License_Server, bytes.NewReader(req_body),
      )
      if err != nil {
         t.Fatal(err)
      }
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         t.Fatal(err)
      }
      res_body, err := io.ReadAll(res.Body)
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      keys, err := mod.GetLicenseKeys(req_body, res_body)
      if err != nil {
         t.Fatal(err)
      }
      if key := hex.EncodeToString(keys[0].Value); key != test.key {
         t.Fatal(key)
      }
      time.Sleep(time.Second)
   }
}
