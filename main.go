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

var configDir string


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

  plex := new(PlexAuto)
  plex = plex.New(args.ConfigDir)

  switch args.Mode {
    case "nzbget":
      plex.NzbGet.Handle(args.NzbGet)

    case "deluge":
      plex.Deluge.Handle(args.Deluge)
  }
}
