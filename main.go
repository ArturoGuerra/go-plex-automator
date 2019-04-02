package main

import (
  "config"
  "handler"
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
  Deluge := deluge.New(&config, FileBot)
  NzbGet := nzbget.New(&config, FileBot)

  return &PlexAuto{
    Config: &config,
    FileBot: FileBot,
    Deluge: Deluge,
    NzbGet: NzbGet,
  }
}

func main () {
  plex := new(PlexAuto)
  plex = plex.New("./config.json")
  h := handler.New()
  args, err := h.Parse()
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  deluge := args.Deluge
  nzbget := args.NzbGet

  NzbGet := plex.NzbGet
  Deluge := plex.Deluge
  switch args.Mode {
    case "nzbget":
      NzbGet.Handle(nzbget)

    case "deluge":
      Deluge.Handle(deluge)
  }
}
