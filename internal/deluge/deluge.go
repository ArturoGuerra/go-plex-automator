package deluge

import (
    "github.com/arturoguerra/goautoplex/internal/filebot"
    "github.com/arturoguerra/goautoplex/internal/structs"
    "github.com/arturoguerra/goautoplex/internal/config"
)


type (
    Deluge struct {
        BaseDir string
        ApiUrl string
        ApiKey string
        Rid string
        FileBot *filebot.FileBot
    }

    RequestBody struct {
        Id int `json:"id"`
        Method string `json:"method"`
        Params []string `json:"params"`
    }
)

func New(c *config.Deluge, f *filebot.FileBot) *Deluge {
    return &Deluge {
        c.BaseDir,
        c.ApiUrl,
        c.ApiKey,
        1,
        f,
    }
}

func (d *Deluge) Handle(args *structs.Deluge) error {
    "TODO"
}
