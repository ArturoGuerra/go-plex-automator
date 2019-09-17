package filebot

import (
    "github.com/arturoguerra/goautoplex/internal/config"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "bytes"

)

type FileBot struct {
    DestDir string
    AmcLogs string
    AmcExclude string
}

func New(c *config.FileBot) *FileBot {
    return &FileBot{
       c.DestDir,
       c.AmcLogs,
       c.AmcExclude,
    }
}

func (f *FileBot) FormatCommand(source string) []string {
    var args []string
    args = append(args, f.DestDir, source, f.AmcLogs, f.AmcExclude)
    return args
}

func (f *FileBot) GetFileBot() string {
    ex, _ := os.Executable()
    return filepath.Dir(ex) + "/filebot"
}

func (f *FileBot) Process(source string) error {
    command := f.FormatCommand(source)
    filebot := f.GetFileBot()
    cmd := exec.Command(filebot, command...)

    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Start()
    if err != nil {
        fmt.Println(err)
        fmt.Println("Error in filebot script")
    }
    cmd.Wait()
    fmt.Println(stderr.String())
    fmt.Println(stdout.String())
    fmt.Println("FileBot Done")
    return nil
}
