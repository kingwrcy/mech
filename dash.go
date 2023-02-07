package mech

import (
   "2a.pages.dev/mech/widevine"
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/mp4"
   "encoding/xml"
   "fmt"
   "io"
   "net/url"
   "os"
)

var client = http.Default_Client

type Stream struct {
   Client_ID string
   Info bool
   Private_Key string
   Poster widevine.Poster
   Name string
   base *url.URL
}

func (s *Stream) DASH(ref string) (dash.Representations, error) {
   res, err := client.Redirect(nil).Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var pres dash.Presentation
   if err := xml.NewDecoder(res.Body).Decode(&pres); err != nil {
      return nil, err
   }
   s.Name = Clean(s.Name)
   s.base = res.Request.URL
   return pres.Representation(), nil
}

func (s Stream) DASH_Get(items dash.Representations, index int) error {
   if s.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(s.Name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("GET", item.Initialization(), nil)
   if err != nil {
      return err
   }
   req.URL = s.base.ResolveReference(req.URL)
   res, err := client.Redirect(nil).Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   media := item.Media()
   pro := http.Progress_Chunks(file, len(media))
   dec := mp4.New_Decrypt(pro)
   var key []byte
   if item.ContentProtection != nil {
      private_key, err := os.ReadFile(s.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(s.Client_ID)
      if err != nil {
         return err
      }
      key_ID, err := widevine.Key_ID(item.ContentProtection.Default_KID)
      if err != nil {
         return err
      }
      mod, err := widevine.New_Module(private_key, client_ID, key_ID)
      if err != nil {
         return err
      }
      keys, err := mod.Post(s.Poster)
      if err != nil {
         return err
      }
      key = keys.Content().Key
      if err := dec.Init(res.Body); err != nil {
         return err
      }
   } else {
      _, err := io.Copy(pro, res.Body)
      if err != nil {
         return err
      }
   }
   for _, ref := range media {
      req, err := http.NewRequest("GET", ref, nil)
      if err != nil {
         return err
      }
      req.URL = s.base.ResolveReference(req.URL)
      res, err := client.Redirect(nil).Level(0).Do(req)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if item.ContentProtection != nil {
         err = dec.Segment(res.Body, key)
      } else {
         _, err = io.Copy(pro, res.Body)
      }
      if err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
