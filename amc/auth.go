package amc

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func (a Auth) Create(name string) error {
   indent, err := json.MarshalIndent(a, "", " ")
   if err != nil {
      return err
   }
   return os.WriteFile(name, indent, os.ModePerm)
}

func (a Auth) Playback(ref string) (*Playback, error) {
   _, nID, found := strings.Cut(ref, "--")
   if !found {
      return nil, fmt.Errorf("%q invalid", ref)
   }
   var p playback_request
   p.Ad_Tags.Mode = "on-demand"
   p.Ad_Tags.URL = "-"
   body, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/playback-id/api/v1/playback/" + nID,
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-Id": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Id": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var play Playback
   if err := json.NewDecoder(res.Body).Decode(&play.body); err != nil {
      return nil, err
   }
   play.head = res.Header
   return &play, nil
}

// this accepts full URL or path only
func (a Auth) Content(raw_ref string) (*Content, error) {
   ref, err := url.Parse(raw_ref)
   if err != nil {
      return nil, err
   }
   var b strings.Builder
   b.WriteString("https://gw.cds.amcn.com")
   b.WriteString("/content-compiler-cr/api/v1/content/amcn/amcplus/path")
   if strings.HasPrefix(ref.Path, "/movies/") {
      b.WriteString("/watch")
   }
   b.WriteString(ref.Path)
   req, err := http.NewRequest("GET", b.String(), nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

type Auth struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func Open_Auth(name string) (*Auth, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   auth := new(Auth)
   if err := json.NewDecoder(file).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func Unauth() (*Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/unauth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Amcn-Device-Id": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(Auth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (a *Auth) Login(email, password string) error {
   buf, err := json.Marshal(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/login",
      bytes.NewReader(buf),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-Id": {"-"},
      "X-Amcn-Device-Id": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Group-Id": {"10"},
      "X-Amcn-Service-Id": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Auth) Refresh() error {
   req, err := http.NewRequest(
      "POST",
      "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/refresh",
      nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("Authorization", "Bearer " + a.Data.Refresh_Token)
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
