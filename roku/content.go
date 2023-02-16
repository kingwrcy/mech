package roku

import (
   "encoding/json"
   "errors"
   "net/url"
   "strings"
   "time"
)

type Content struct {
   Meta struct {
      ID string
      MediaType string
   }
   Title string
   Series struct {
      Title string
   }
   SeasonNumber string
   EpisodeNumber string
   ReleaseDate string
   RunTimeSeconds int64
   ViewOptions []struct {
      License string
      Media struct {
         Videos []Video
      }
   }
}

func (c Content) Name() string {
   var buf strings.Builder
   if c.Meta.MediaType == "episode" {
      buf.WriteString(c.Series.Title)
      buf.WriteByte('-')
      buf.WriteString(c.SeasonNumber)
      buf.WriteByte('-')
      buf.WriteString(c.EpisodeNumber)
      buf.WriteByte('-')
   }
   buf.WriteString(c.Title)
   return buf.String()
}

func (c Content) String() string {
   var buf strings.Builder
   write := func(str string) {
      buf.WriteString(str)
   }
   write("ID: ")
   write(c.Meta.ID)
   write("\nType: ")
   write(c.Meta.MediaType)
   write("\nTitle: ")
   write(c.Title)
   if c.Meta.MediaType == "episode" {
      write("\nSeries: ")
      write(c.Series.Title)
      write("\nSeason: ")
      write(c.SeasonNumber)
      write("\nEpisode: ")
      write(c.EpisodeNumber)
   }
   write("\nDate: ")
   write(c.ReleaseDate)
   write("\nDuration: ")
   write(c.Duration().String())
   return buf.String()
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
   return time.Duration(c.RunTimeSeconds) * time.Second
}

func (c Content) DASH() *Video {
   for _, opt := range c.ViewOptions {
      for _, vid := range opt.Media.Videos {
         if vid.VideoType == "DASH" {
            return &vid
         }
      }
   }
   return nil
}

func (c Content) HLS() (*Video, error) {
   for _, opt := range c.ViewOptions {
      for _, vid := range opt.Media.Videos {
         if vid.DrmAuthentication == nil {
            if vid.VideoType == "HLS" {
               return &vid, nil
            }
         }
      }
   }
   return nil, errors.New("drmAuthentication")
}
