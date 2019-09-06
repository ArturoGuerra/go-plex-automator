package main

import (
  "goplex/internal/config"
  "goplex/internal/parameters"
  "goplex/internal/filebot"
  "goplex/internal/deluge"
  "goplex/internal/nzbget"
  "goplex/internal/fileloader"
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

func New(configFilename string) *PlexAuto {
  config, err := fileloader.LoadConfig(configFilename)
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
  args, err := parameters.Parse()
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  if configDir == "" {
      configDir = "./plexbot.conf"
  }

  plex := New(configDir)

  switch args.Mode {
    case "nzbget":
      fmt.Println("Running in nzbget mode")
      plex.NzbGet.Handle(args.NzbGet)

    case "deluge":
      fmt.Println("Running in deluge mode")
      plex.Deluge.Handle(args.Deluge)
  }
}
