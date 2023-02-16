package roku

import (
   "encoding/json"
   "errors"
   "net/url"
   "strings"
   "time"
)

func (c Content) Name() string {
   var b strings.Builder
   if c.Meta.Media_Type == "episode" {
      b.WriteString(c.Series.Title)
      b.WriteByte('-')
      b.WriteString(c.Season_Number)
      b.WriteByte('-')
      b.WriteString(c.Episode_Number)
      b.WriteByte('-')
   }
   b.WriteString(c.Title)
   b.WriteByte('-')
   b.WriteString(c.Year())
   return b.String()
}

func (c Content) Year() string {
   year, _, _ := strings.Cut(c.Release_Date, "-")
   return year
}

type Content struct {
   Episode_Number string `json:"episodeNumber"`
   Meta struct {
      ID string
      Media_Type string `json:"mediaType"`
   }
   Release_Date string `json:"releaseDate"` // 2007-01-01T000000Z
   Run_Time_Seconds int64 `json:"runTimeSeconds"`
   Season_Number string `json:"seasonNumber"`
   Series struct {
      Title string
   }
   Title string
   View_Options []struct {
      License string
      Media struct {
         Videos []Video
      }
   } `json:"viewOptions"`
}

func (c Content) String() string {
   var b strings.Builder
   b.WriteString("ID: ")
   b.WriteString(c.Meta.ID)
   b.WriteString("\nType: ")
   b.WriteString(c.Meta.Media_Type)
   b.WriteString("\nTitle: ")
   b.WriteString(c.Title)
   if c.Meta.Media_Type == "episode" {
      b.WriteString("\nSeries: ")
      b.WriteString(c.Series.Title)
      b.WriteString("\nSeason: ")
      b.WriteString(c.Season_Number)
      b.WriteString("\nEpisode: ")
      b.WriteString(c.Episode_Number)
   }
   b.WriteString("\nDate: ")
   b.WriteString(c.Release_Date)
   b.WriteString("\nDuration: ")
   b.WriteString(c.Duration().String())
   return b.String()
}

func New_Content(id string) (*Content, error) {
   var ref url.URL
   ref.Scheme = "https"
   ref.Host = "content.sr.roku.com"
   ref.Path = "/content/v1/roku-trc/" + id
   ref.RawQuery = url.Values{
      "expand": {"series"},
      "include": {strings.Join([]string{
         "episodeNumber",
         "releaseDate",
         "runTimeSeconds",
         "seasonNumber",
         // this needs to be exactly as is, otherwise size blows up
         "series.seasons.episodes.viewOptions\u2008",
         "series.title",
         "title",
         "viewOptions",
      }, ",")},
   }.Encode()
   var buf strings.Builder
   buf.WriteString("https://therokuchannel.roku.com/api/v2/homescreen/content/")
   buf.WriteString(url.PathEscape(ref.String()))
   res, err := Client.Get(buf.String())
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

func (c Content) Duration() time.Duration {
   return time.Duration(c.Run_Time_Seconds) * time.Second
}

func (c Content) DASH() *Video {
   for _, opt := range c.View_Options {
      for _, vid := range opt.Media.Videos {
         if vid.Video_Type == "DASH" {
            return &vid
         }
      }
   }
   return nil
}

func (c Content) HLS() (*Video, error) {
   for _, opt := range c.View_Options {
      for _, vid := range opt.Media.Videos {
         if vid.DRM_Authentication == nil {
            if vid.Video_Type == "HLS" {
               return &vid, nil
            }
         }
      }
   }
   return nil, errors.New("drmAuthentication")
}
