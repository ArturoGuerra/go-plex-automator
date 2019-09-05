package deluge

import (
  "net/http"
  "goplex/src/filebot"
  "encoding/json"
  "goplex/src/config"
  "errors"
  "bytes"
  "goplex/src/utils"
  "fmt"
)

type (
  Deluge struct {
    Password string
    Api string
    TorrentsDir string
    Rid int
    FileBot *filebot.FileBot
  }

  RequestBody struct {
    Id int `json:"id"`
    Method string `json:"method"`
    Params []string `json:"params"`
  }
)

func New(c *config.Configuration, f *filebot.FileBot) *Deluge {
  dc := c.Deluge
  return &Deluge{
    dc.Password,
    dc.Api,
    dc.TorrentsDir,
    1,
    f,
  }
}

func (d *Deluge) Handle(args *utils.Deluge) error {
  var mode string
  var err error
  pass := false
  filebot := d.FileBot


  switch args.TorrentDir {
  case d.TorrentsDir + "/Shows":
    mode = "shows"
    pass = true
    fmt.Println("Shows deluge dir")
  case d.TorrentsDir + "/Movies":
    mode = "movies"
    fmt.Println("Movies deluge dir")
    pass = true
  default:
    err = errors.New("Invalid torrent dir")
    fmt.Println("Invalid deluge dir")
  }

  if pass {
    if err = filebot.Handle(mode, d.TorrentsDir); err != nil {
      return err
    }
  }

  err = d.Clean(args.TorrentId)
  if err != nil {
      fmt.Println("Error cleaning deluge: ", err)
      return err
  }

  return nil
}

func (d *Deluge) Clean(id string) error {
  fmt.Println("Starting deluge clean for: ", id)
  client := http.Client{}
  buf := new(bytes.Buffer)
  body := &RequestBody{
    d.Rid,
    "auth.login",
    []string{d.Password},
  }
  json.NewEncoder(buf).Encode(body)
  resp, err := client.Post(d.Api, "application/json", buf)
  if err != nil {
    return err
  }

  d.Rid = d.Rid + 1

  buf = new(bytes.Buffer)
  body = &RequestBody{
    d.Rid,
    "webapi.remove_torrent",
    []string{id,"true"},
  }
  json.NewEncoder(buf).Encode(body)
  req, err := http.NewRequest("POST", d.Api, buf)
  req.Header.Set("Content-Type", "application/json")

  for _, cookie := range resp.Cookies() {
      req.AddCookie(cookie)
  }

  resp, err = client.Do(req)
  if err != nil {
    return err
  }

  if resp.StatusCode == 200 {
      fmt.Println("Deluge clean completed successfully")
  } else {
      fmt.Println("There was an error cleaning deluge, please clean the client manually")
  }

  return nil
}
