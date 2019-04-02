package filebot

import (
  "config"
  "os"
  "errors"
  "bytes"
  "os"
  "utils"
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

func (f FileBot) New(c *config.Configuration) *FileBot {
  return &FileBot{
    c.Plex,
    c.SickRage.Api,
    c.SickRage.Url,
    c.CouchPotato.Api,
    c.CouchPotato.Url,
    c.FileBot.Logs,
    c.FileBot.Amc,
  }
}


func (f *FileBot) Movies(root string) error {
  base := root + "/" + "Movies"
  callback := "curl " + f.CouchPotatoUrl + "/api/" + f.CouchPotatoApi + "/manage.update"
  err := f.Process(source, callback)
  if err != nil {
    return err
  }

  return nil
}

func (f *FileBot) Shows(root string) error {
  base := root + "/" + "Shows"
  callback := "curl " + f.SickRageUrl + "/api/" + f.SickRageApi + "?cmd=show.refresh&tvdbid={info.id}"
  err := f.Process(source, callback)
  if err != nil {
    return err
  }

  return nil
}

func (f *FileBot) Anime(root string) error {
  source := root + "/" + "Anime"
  callback := "curl " + f.SickRageUrl + "/api/" + f.SickRageApi + "?cmd=show.refresh&tvdbid={info.id}"
  err := f.Process(source, callback)
  if err != nil {
    return err
  }

  return nil

}

func FormatCommand (source, callback string) string {
  var buf bytes.Buffer
  buf.WriteString("/usr/bin/filebot -script fn:amc --output ")
  buf.WriteString(f.DestinationDir)
  buf.WriteString(" --action copy -non-strict ")
  buf.WriteString(source)
  buf.WriteString(" --conflict override --log-file ")
  buf.WriteString(f.AmcLogs)
  buf.WriteString(" --def subtitles=en,es --def excludeList=")
  buf.WriteString(f.AmcExclude)
  buf.WriteString(" --def clean=y --def unsorted=y --def extras=y --def seriesFormat=")
  buf.WriteString("\"" + f.DestinationDir + "/TV Shows/{n.replaceAll(/'/)}/Season {s.pad(2)}/{n} - {s00e00} - {t}\"")
  buf.WriteString(" --def animeFormat=")
  buf.WriteString("\"" + f.DestinationDir + "/Anime/{n.replaceAll(/'/)}/Season {s.pad(2)}/{n} - {s00e00} - {t}\"")
  buf.WriteString(" --def exec=\"" + callback + "\"")
  buf.WriteString(" --def minLengthMS=300000")
  result := buf.String()
  return result
}

func (f *FileBot) Process(source, callback string) error {
  command := FormatCommand(source, callback)
  cmd := exec.Command(command)
  out, err := cmd.CombinedOutput()
  if err != nil {
    return err
  }

  fmt.Println(string(out))
  return nil
}

func Handle(f *FileBot) (mode, source string) error {
  err := nil
  switch mode {
  case "movies":
    err = f.Movies(source)
  case "shows":
    err = f.Shows(source)
  case "anime":
    err = f.Anime(source)
  }

  return err
}
