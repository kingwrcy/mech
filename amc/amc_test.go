package amc

import (
   "2a.pages.dev/mech/widevine"
   "fmt"
   "os"
   "testing"
   "time"
)

var links = []string{
   // amcplus.com/movies/jerry-maguire--1054053
   "/movies/jerry-maguire--1054053",
   // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
   "/shows/orphan-black/episodes/season-1-instinct--1011152",
}

func Test_Content(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Open_Auth(home + "/mech/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, link := range links {
      con, err := auth.Content(link)
      if err != nil {
         t.Fatal(err)
      }
      video, err := con.Video_Player()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(video.Name())
      time.Sleep(time.Second)
   }
}

// amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
const (
   key = "a66a5603545ad206c1a78e160a6710b1"
   nID = 1011152
   raw_key_ID = "c0e598b247fa443590299d5ef47da32c"
)

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
   key_ID, err := widevine.Key_ID(raw_key_ID)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Open_Auth(home + "/mech/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback(nID)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(play)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != key {
      t.Fatal(keys)
   }
}
