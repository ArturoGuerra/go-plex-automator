package nzbget

import (
  "goplex/src/filebot"
  "goplex/src/config"
  "goplex/src/utils"
  "strconv"
  "errors"
  "os"
  "fmt"
)

type (
  NzbGet struct {
    BaseDir string
    FileBot *filebot.FileBot
    Success int
    Error int
  }
)

func New(c *config.Configuration, f *filebot.FileBot) *NzbGet {
  nc := &c.NzbGet
  return &NzbGet{
    nc.BaseDir,
    f,
    93,
    94,
  }
}

func Clean(dir string) {
    fmt.Println("Removing Nzb File: ", dir)
    os.RemoveAll(dir)
}

func (n *NzbGet) Handle(args *utils.NzbGet) error {
  filebot := n.FileBot
  var err error

  switch args.NzbppTotalStatus {
  case "SUCCESS":
    fmt.Println("Nzbget succeded at processing nzb, starting post processing")
    err = filebot.Handle(args.NzbppCategory, n.BaseDir)
    Clean(args.NzbppDirectory)
    if err != nil {
      os.Exit(n.Error)
      return err
    }

    os.Exit(n.Success)
    return nil

  case "FAILURE":
    fmt.Println("Nzb failed, ignoring and cleaning up")
    Clean(args.NzbppDirectory)
    os.Exit(n.Error)
    err = errors.New(strconv.Itoa(n.Error))
    return err

  default:
    fmt.Println("NzbGet returned an unknown error code")
    Clean(args.NzbppDirectory)
    os.Exit(n.Error)
    err = errors.New(strconv.Itoa(n.Error))
    return err
  }

  return nil
}
