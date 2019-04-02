package main

import (
  "config"
  "filebot"
  "deluge"
  "nzbget"
  "os"
  "fmt"
)

type PlexAuto struct {
  Config *config.Configuration
  FileBot *filebot.FileBot
  Deluge *deluge.Deluge
  NzbGet *nzbget.NzbGet

}

func (p PlexAuto) New(configFilename string) *PlexAuto {
  config, err := config.LoadConfig(configFilename)
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  FileBot := filebot.New(&config)
  Deluge := deluge.New(&config, &FileBot)
  NzbGet := nzbget.New(&config, &FileBot)

  return &PlexAuto{
    Config: &config,
    Filebot: &FileBot,
    Deluge: &Deluge,
    NzbGet: &NzbGet,
  }
}

func main () {
  plex := PlexAuto.New("./config.json")
  handler := Handler.New()
  args, err := handler.Parse()
  if err != nil {
    os.Exit(3)
  }

  if err != nil {
    os.Exit(3)
  }

  deluge := &args.Deluge
  nzbget := &args.NzbGet

  switch args.Mode {
    case "nzbget":
      plex.NzbGet(nzbget)

    case "deluge":
      plex.Deluge(deluge)
  }
}
