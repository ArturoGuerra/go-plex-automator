package api

import (
    "os"
    "net/http"
    "github.com/labstack/echo"
    "github.com/arturoguerra/goautoplex/internal/structs"
    "github.com/arturoguerra/goautoplex/internal/filebot"
    "github.com/arturoguerra/goautoplex/internal/nzbget"
    "github.com/arturoguerra/goautoplex/internal/deluge"
    "github.com/arturoguerra/goautoplex/internal/config"
)

func LoadCfg () (*config.FileBot, *config.Deluge, *config.NzbGet) {

    return filebot, deluge, nzbget
}

type Api struct {
    Deluge *deluge.Deluge
    NzbGet *nzbget.NzbGet
}

func New() *Api {
    fbcfg, dcfg, ngcfg := LoadCfg()

    Filebot := filebot.New(fbcfg)
    Deluge := deluge.New(dcfg, Filebot)
    NzbGet := nzbget.New(ngcfg, Filebot)

    return &Api{
        Deluge,
        NzbGet,
    }
}

func (a *Api) NzbGetHandler(c echo.Context) error {
}

func (a *Api) DelugeHandler(c echo.Context) error {
}
