package nzbget

import (
    "github.com/arturoguerra/goautoplex/internal/filebot"
    "github.com/arturoguerra/goautoplex/internal/structs"
    "github.com/arturoguerra/goautoplex/internal/config"
    "github.com/arturoguerra/goautoplex/internal/datavalidator/revars"
    "errors"
    "fmt"
    "os"
)

type NzbGet struct {
    BaseDir string
    FileBot *filebot.FileBot
}

var Error int = 94
var Success int = 93

func New(c *config.NzbGet, f *filebot.FileBot) *NzbGet {
    return &NzbGet{
        c.BaseDir,
        f,
    }
}

func clean(dir string) {
    fmt.Println("Removing Nzb File:", dir)
    os.RemoveAll(dir)
}

func getnzbname(dir string) string {
    return revars.LinuxPath.FindStringSubmatch(dir)[1]
}

func (n *NzbGet) Handle(payload *structs.NzbGetPayload) (err error) {
    switch payload.Status {
    case "SUCCESS":
        fmt.Println("Processing NzbGet...")
        targetName := getnzbname(payload.Dir)
        source := n.BaseDir + "/" + payload.Category + "/" + targetName

        err = n.FileBot.Process(source)
        if err != nil {
            return err
        }

        clean(payload.Dir)

        return nil

    case "FAILURE":
        e := "NzbGet failure, stopping post processing..."
        fmt.Println(e)
        err = errors.New(e)
        return err

    default:
        e := "NzbGet error processing, unknown error"
        fmt.Println(e)
        err = errors.New(e)
        return err
    }
}
