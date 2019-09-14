package filebot

import (
    "fmt"
)

type FileBot struct {
    DestDir string
    AmcLogs string
    AmcExclude string
}

func New(c *config.FileBot) *FileBot {
    return {
       c.DestDir,
       c.AmcLogs,
       c.AmcExclude,
   }
}

func (f *FileBot) Process("TODO") error {
}
