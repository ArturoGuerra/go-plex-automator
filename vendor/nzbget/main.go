package nzbget

import (
  "filebot"
  "encoding/json"
  "config"
  "utils"
  "errors"
  "os"
)

type (
  NzbGet struct {
    BaseDir string
    FileBot *filebot.FileBot
    Success int
    Error int
  }
)

func (n NzbGet) New(c *config.Configuration, f *filebot.FileBot) *NzbGet {
  nc := &c.NzbGet
  return &NzbGet{
    *nc.BaseDir,
    &f,
    93,
    94,

  }
}

func (n *NzbGet) Handle(args *utils.NzbGet) error {
  filebot := n.FileBot
  var err error

  switch args.NzbppTotalStatus {
  case "SUCCESS":
    err = filebot.Handle(args.NzbppCategory, n.BaseDir)
    if err != nil {
      os.Exit(e.Error)
      return err
    }

    os.Exit(e.Success)
    return nil

  case "FAILURE":
    os.Exit(n.Error)
    err = errors.New(n.Error)
    return err

  default:
    os.Exit(n.Error)
    err = errors.New(n.Error)
    return err
  }

  return nil
}

func
