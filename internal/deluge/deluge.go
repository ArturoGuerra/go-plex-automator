package deluge

import (
    "github.com/arturoguerra/goautoplex/internal/datavalidator/revars"
    "github.com/arturoguerra/goautoplex/internal/filebot"
    "github.com/arturoguerra/goautoplex/internal/structs"
    "github.com/arturoguerra/goautoplex/internal/config"
    "errors"
    "fmt"
)


type (
    Deluge struct {
        BaseDir string
        ApiUrl string
        ApiKey string
        Rid int
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

// direct download dir specified in a given label
func getlabel(labelDir string) (error, string) {
    labels := []string{"movies", "shows"}
    label := revars.LinuxPath.FindStringSubmatch(labelDir)[1]

    for _, l := range labels {
        if l == label {
            return nil, label
        }
    }

    e := "Invalid deluge label: " + label
    fmt.Println(e)
    err := errors.New(e)
    return err, label
}


func (d *Deluge) Handle(payload *structs.DelugePayload) (err error) {
    err, label := getlabel(payload.Dir)
    if err != nil {
        return err
    }

    fmt.Println("Processing deluge with id: ", payload.Id)
    source := d.BaseDir + "/" + label + "/" + payload.Name

    err = d.FileBot.Process(source)
    if err != nil {
        return err
    }

    return nil
}
