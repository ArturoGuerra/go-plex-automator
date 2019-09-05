package filebot

import (
  "goplex/config"
  "bytes"
  "errors"
  "os/exec"
  "fmt"
  "os"
  "path/filepath"
)

type (
  FileBot struct {
    DestinationDir string
    SickRageApi string
    SickRageUrl string
    CouchPotatoApi string
    CouchPotatoUrl string
    AmcLogs string
    AmcExclude string
  }
)

func New(c *config.Configuration) *FileBot {
  csr := c.SickRage
  ccp := c.CouchPotato
  cfb := c.FileBot
  return &FileBot{
    c.Plex,
    csr.Api,
    csr.Url,
    ccp.Api,
    ccp.Url,
    cfb.Logs,
    cfb.Amc,
  }
}


func (f *FileBot) Movies(root string) {
  callback := "curl " + f.CouchPotatoUrl + "/api/" + f.CouchPotatoApi + "/manage.update"
  root = root + "/Movies"
  f.Process(root, callback)
}

func (f *FileBot) Shows(root string) {
  callback := "curl " + f.SickRageUrl + "/api/" + f.SickRageApi + "?cmd=show.refresh&tvdbid={info.id}"
  root = root + "/Shows"
  f.Process(root, callback)
}

func (f *FileBot) Anime(root string) {
  callback := "curl " + f.SickRageUrl + "/api/" + f.SickRageApi + "?cmd=show.refresh&tvdbid={info.id}"
  root = root + "/Anime"
  f.Process(root, callback)

}

func FormatCommand(source, callback string, f *FileBot) []string {
    var args []string
    args = append(args, f.DestinationDir, source, f.AmcLogs, f.AmcExclude, f.DestinationDir, callback)
    return args
}

func getFileBot() string {
    ex, _ := os.Executable()
    return filepath.Dir(ex) + "/filebot"
}


func (f *FileBot) Process(source, callback string) {
  command := FormatCommand(source, callback, f)
  fmt.Println("Running FileBot...")
  filebot := getFileBot()
  fmt.Println(filebot, command)
  cmd := exec.Command(filebot, command...)
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr

  err := cmd.Start()
  if err != nil {
      fmt.Println(err)
      fmt.Println("Error in filebot")
  }
  cmd.Wait()

  fmt.Println(stderr.String())
  fmt.Println(out.String())
  fmt.Println("FileBot Done")
}

func (f *FileBot) Handle(mode, source string) error {
  var err error
  switch mode {
  case "movies":
    f.Movies(source)
  case "shows":
    f.Shows(source)
  case "tv":
    f.Shows(source)
  case "anime":
    f.Anime(source)
  default:
    err = errors.New("Invalid Mode")
  }


  return err
}
