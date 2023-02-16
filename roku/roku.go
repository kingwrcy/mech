package roku

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "bytes"
   "io"
)

type Playback struct {
   DRM struct {
      Widevine struct {
         License_Server string `json:"licenseServer"`
      }
   }
}

func (p Playback) Request_URL() string {
   return p.DRM.Widevine.License_Server
}

func (Playback) Request_Header() http.Header {
   return nil
}

func (Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

var Client = http.Default_Client

type Cross_Site struct {
   cookie *http.Cookie // has own String method
   token string
}

func (c Cross_Site) Playback(id string) (*Playback, error) {
   buf, err := json.Marshal(map[string]string{
      "mediaFormat": "mpeg-dash",
      "rokuId": id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://therokuchannel.roku.com/api/v3/playback",
      bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
   }
   req.AddCookie(c.cookie)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Video struct {
   DRM_Authentication *struct{} `json:"drmAuthentication"`
   URL string
   Video_Type string `json:"videoType"`
}

func New_Cross_Site() (*Cross_Site, error) {
   // this has smaller body than www.roku.com
   res, err := Client.Get("https://therokuchannel.roku.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var site Cross_Site
   for _, cook := range res.Cookies() {
      if cook.Name == "_csrf" {
         site.cookie = cook
      }
   }
   var scan json.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte("\tcsrf:")
   scan.Scan()
   scan.Sep = nil
   if err := scan.Decode(&site.token); err != nil {
      return nil, err
   }
   return &site, nil
}
