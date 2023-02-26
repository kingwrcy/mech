package roku

import (
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

const (
   episode = iota
   movie
)

type key struct {
   media_type int
   content_ID bool
}

var tests = map[key]struct {
   key string
   playback_ID string
   pssh string
} {
   // therokuchannel.roku.com/watch/12cdc6139bb858f4a597a1aa51a1e37d
   {episode, false}: {
      key: "17e71e939c2a541142112c376cf73cc4",
      playback_ID: "12cdc6139bb858f4a597a1aa51a1e37d",
      pssh: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQmKI1QsJl+5Fx3Kpnu/oGTBoKaW50ZXJ0cnVzdCIBKg==",
   },
   // therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
   {episode, true}: {
      key: "e258b67d75420066c8424bd142f84565",
      playback_ID: "105c41ea75775968b670fbb26978ed76",
      pssh: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg==",
   },
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   {movie, false}: {
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      playback_ID: "597a64a4a25c5bf6af4a8c7053049a6f",
   },
   // therokuchannel.roku.com/watch/2b3166271d83569c81d41030e9ba7fb0
   {movie, true}: {playback_ID: "2b3166271d83569c81d41030e9ba7fb0"},
}

func Test_Video(t *testing.T) {
   test := tests[key{episode, false}]
   con, err := New_Content(test.playback_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(con)
   fmt.Printf("%+v\n", con.DASH())
}

func Test_Playback(t *testing.T) {
   site, err := New_Cross_Site()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := site.Playback(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", play)
   }
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
      if test.pssh != "" {
         pssh, err := base64.StdEncoding.DecodeString(test.pssh)
         if err != nil {
            t.Fatal(err)
         }
         mod, err := widevine.New_Module(private_key, client_ID, pssh)
         if err != nil {
            t.Fatal(err)
         }
         site, err := New_Cross_Site()
         if err != nil {
            t.Fatal(err)
         }
         play, err := site.Playback(test.playback_ID)
         if err != nil {
            t.Fatal(err)
         }
         keys, err := mod.Post(play)
         if err != nil {
            t.Fatal(err)
         }
         if keys.Content().String() != test.key {
            t.Fatal(keys)
         }
      }
   }
}
