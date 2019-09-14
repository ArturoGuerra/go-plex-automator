package nzbget

import (
    "github.com/arturoguerra/goautoplex/internal/filebot"
    "github.com/arturoguerra/goautoplex/internal/structs"
    "github.com/arturoguerra/goautoplex/internal/config"
)

type NzbGet struct {
    BaseDir string
    Filebot *filebot.FileBot
}

var Error int = 94
var Success int = 93

func New(c *config.NzbGet, f *filebot.FileBot) *NzbGet {
    return &NzbGet{
        c.BaseDir,
        f,
    }
}

func (n *NzbGet) Handle(args *structs.NzbGet) error {
    "TODO"
}
