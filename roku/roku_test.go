package roku

import (
   "2a.pages.dev/mech/widevine"
   "fmt"
   "os"
   "testing"
)

var tests = []struct {
   key string
   playback_ID string
   raw_key_ID string
}{
   {
      // episode
      // therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
      playback_ID: "105c41ea75775968b670fbb26978ed76",
   }, {
      // therokuchannel.roku.com/watch/2b3166271d83569c81d41030e9ba7fb0
      playback_ID: "2b3166271d83569c81d41030e9ba7fb0",
   }, {
      // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      playback_ID: "597a64a4a25c5bf6af4a8c7053049a6f",
      raw_key_ID: "28339AD78F734520DA24E6E0573D392E",
   },
}

func Test_Video(t *testing.T) {
   con, err := New_Content(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(con)
   video, err := con.HLS()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
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
      if test.key != "" {
         key_ID, err := widevine.Key_ID(test.raw_key_ID)
         if err != nil {
            t.Fatal(err)
         }
         mod, err := widevine.New_Module(private_key, client_ID, key_ID)
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
