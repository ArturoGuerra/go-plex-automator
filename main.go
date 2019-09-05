package main

import (
  "goplex/config"
  "goplex/internal/handler"
  "goplex/internal/filebot"
  "goplex/internal/deluge"
  "goplex/internal/nzbget"
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
  if configDir == "" {
      configDir = "./plexbot.conf"
  }

  plex = plex.New(configDir)

  switch args.Mode {
    case "nzbget":
      fmt.Println("Running in nzbget mode")
      plex.NzbGet.Handle(args.NzbGet)

    case "deluge":
      fmt.Println("Running in deluge mode")
      plex.Deluge.Handle(args.Deluge)
  }
}
