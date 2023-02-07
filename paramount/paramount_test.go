package paramount

import (
   "fmt"
   "testing"
   "time"
)

var tests = map[test_type]string{
   // cbs.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_
   {episode, hls_clear}: "YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_",
   // paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD
   {movie, dash_cenc}: "tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD",
   // paramountplus.com/shows/video/622520382
   {episode, stream_pack}: "622520382",
   // paramountplus.com/movies/video/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
   {movie, stream_pack}: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
}

type test_type struct {
   content_type int
   asset int
}

const (
   dash_cenc = iota
   episode
   hls_clear
   movie
   stream_pack
)

func Test_Preview(t *testing.T) {
   for _, test := range tests {
      preview, err := New_Preview(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", preview)
      time.Sleep(time.Second)
   }
}
