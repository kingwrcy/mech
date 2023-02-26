package widevine

import (
   "2a.pages.dev/mech/roku"
   "bytes"
   "encoding/base64"
   "fmt"
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
   // therokuchannel.roku.com/watch/2b3166271d83569c81d41030e9ba7fb0
   {movie, 500}: {playback_ID: "2b3166271d83569c81d41030e9ba7fb0"},
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   {movie, 200}: {
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      playback_ID: "597a64a4a25c5bf6af4a8c7053049a6f",
   },
   // therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
   {episode, 500}: {
      init_data: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg==",
      key: "e258b67d75420066c8424bd142f84565",
      playback_ID: "105c41ea75775968b670fbb26978ed76",
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
      if test.init_data != "" {
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
         fmt.Println(keys)
         time.Sleep(time.Second)
      }
   }
}
