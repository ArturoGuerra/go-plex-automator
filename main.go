package main

import (
  "config"
  "handler"
  "filebot"
  "deluge"
  "nzbget"
  "os"
  "fmt"
  "flag"
)

var configDir string

func init () {
    flag.StringVar(&configDir, "configdir", "configdir", "Configuration file dir")
}

type PlexAuto struct {
  Config *config.Configuration
  FileBot *filebot.FileBot
  Deluge *deluge.Deluge
  NzbGet *nzbget.NzbGet
  configDir string

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
  plex = plex.New(configDir)
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

  switch args.Mode {
    case "nzbget":
      plex.NzbGet.Handle(args.NzbGet)

    case "deluge":
      plex.Deluge.Handle(args.Deluge)
  }
}
