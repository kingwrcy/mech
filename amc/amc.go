package amc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "strings"
)

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

type playback_request struct {
   Ad_Tags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      Player_Height int `json:"playerHeight"`
      Player_Width int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

func (Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (p Playback) Request_Header() http.Header {
   token := p.head.Get("X-AMCN-BC-JWT")
   head := make(http.Header)
   head.Set("bcov-auth", token)
   return head
}

type Playback struct {
   head http.Header
   body struct {
      Data struct {
         Playback_JSON_Data struct {
            Sources []Source
         } `json:"playbackJsonData"`
      }
   }
}

func (p Playback) Source() *Source {
   for _, item := range p.body.Data.Playback_JSON_Data.Sources {
      if item.Type == "application/dash+xml" {
         return &item
      }
   }
   return nil
}

func (p Playback) Request_URL() string {
   return p.Source().Key_Systems.Widevine.License_URL
}

func (v Video_Player) Name() string {
   var b strings.Builder
   b.WriteString(v.Current_Video.Meta.Title)
   b.WriteByte('-')
   b.WriteString(v.Year())
   return b.String()
}

type Video_Player struct {
   Current_Video struct {
      Meta struct {
         Airdate string // 1996-01-01T00:00:00.000Z
         Title string
      }
   } `json:"currentVideo"`
}

func (v Video_Player) Year() string {
   year, _, _ := strings.Cut(v.Current_Video.Meta.Airdate, "-")
   return year
}

func (c Content) Video_Player() (*Video_Player, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         vid := new(Video_Player)
         err := json.Unmarshal(child.Properties, vid)
         if err != nil {
            return nil, err
         }
         return vid, nil
      }
   }
   return nil, errors.New("video-player-ap not present")
}

var Client = http.Default_Client
