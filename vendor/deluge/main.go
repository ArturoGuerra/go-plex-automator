package deluge

import (
  "net/http"
  "filebot"
  "encoding/json"
  "config"
  "bytes"
  "utils"
)

type (
  Deluge struct {
    Port string
    Host string
    User string
    Password string
    Api string
    TorrentsDir string
    Rid int
    FileBot *filebot.FileBot
  }

  RequestBody struct {
    Id: int `json:"id"`
    Method string `json:"method"`
    Params []string `json:"params"`
  }
)

func (d Deluge) New(c *config.Configuration, f *filebot.FileBot) *Deluge {
  return &Deluge{
    c.Port,
    c.Host,
    c.User,
    c.Password,
    c.Api,
    c.TorrentsDir,
    1,
    &f,
}

func (d *Deluge) Handle(a *utils.Deluge) error {

}

func (d *Deluge) Clean(id int) error {
  client := http.Client{}
  buf := new(bytes.Buffer)
  body := &RequestBody{
    c.Rid,
    "auth.login",
    []string{c.Password},
  }
  json.NewEncoder(buf).Encode(body)
  resp, err := client.Post(d.Api, "application/json", buf)
  if err != nil {
    return err
  }

  c.Rid = c.Rid + 1

  client = http.Client{
    Jar: resp.Cookies(),
  }

  buf = new(bytes.Buffer)
  body = &RequestBody{
    c.Rid,
    "webapi.remove_torrent",
    []string{id,true}
  }
  json.NewEncoder(buf).Encode(body)
  _, err = client.Post(d.Api, "application/json", buf)
  if err != nil {
    return err
  }
}
