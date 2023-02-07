package mech

import (
   "2a.pages.dev/rosso/hls"
   "2a.pages.dev/rosso/http"
   "fmt"
   "io"
   "os"
)

func (s Stream) HLS_Streams(items hls.Streams, index int) error {
   return hls_get(s, items, index)
}

func (s Stream) HLS_Media(items hls.Media, index int) error {
   return hls_get(s, items, index)
}

func (s *Stream) HLS(ref string) (*hls.Master, error) {
   res, err := client.Redirect(nil).Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s.Name = Clean(s.Name)
   s.base = res.Request.URL
   return hls.New_Scanner(res.Body).Master()
}

func hls_get[T hls.Mixed](str Stream, items []T, index int) error {
   if str.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(str.Name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("GET", item.URI(), nil)
   if err != nil {
      return err
   }
   req.URL = str.base.ResolveReference(req.URL)
   res, err := client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.New_Scanner(res.Body).Segment()
   if err != nil {
      return err
   }
   var block *hls.Block
   if seg.Key != "" {
      res, err := client.Get(seg.Key)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      key, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      block, err = hls.New_Block(key)
      if err != nil {
         return err
      }
   }
   pro := http.Progress_Chunks(file, len(seg.URI))
   for _, ref := range seg.URI {
      req, err := http.NewRequest("GET", ref, nil)
      if err != nil {
         return err
      }
      req.URL = res.Request.URL.ResolveReference(req.URL)
      res, err := client.Level(0).Redirect(nil).Do(req)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if block != nil {
         text, err := io.ReadAll(res.Body)
         if err != nil {
            return err
         }
         text = block.Decrypt_Key(text)
         if _, err := pro.Write(text); err != nil {
            return err
         }
      } else {
         _, err := io.Copy(pro, res.Body)
         if err != nil {
            return err
         }
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
