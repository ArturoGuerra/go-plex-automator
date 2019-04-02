package deluge

import (
  "net/http"
  "filebot"
  "encoding/json"
  "config"
  "errors"
  "bytes"
  "utils"
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
  filebot := d.FileBot

  switch args.TorrentDir {
  case d.TorrentsDir + "/Shows":
    mode = "shows"
  case d.TorrentsDir + "/Movies":
    mode = "movies"
  default:
    err = errors.New("Invalid torrent dir")
    return err
  }

  err = filebot.Handle(mode, d.TorrentsDir)
  if err != nil {
    return err
  }

  err = d.Clean(args.TorrentId)
  if err != nil {
      return err
  }

  return nil
}

func (d *Deluge) Clean(id string) error {
  client := http.Client{}
  buf := new(bytes.Buffer)
  body := &RequestBody{
    d.Rid,
    "auth.login",
    []string{d.Password},
  }
  json.NewEncoder(buf).Encode(body)
  _, err := client.Post(d.Api, "application/json", buf)
  if err != nil {
    return err
  }

  d.Rid = d.Rid + 1

  client = http.Client{}

  buf = new(bytes.Buffer)
  body = &RequestBody{
    d.Rid,
    "webapi.remove_torrent",
    []string{id,"true"},
  }
  json.NewEncoder(buf).Encode(body)
  _, err = client.Post(d.Api, "application/json", buf)
  if err != nil {
    return err
  }

  return nil
}
