package api

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/arturoguerra/goautoplex/internal/structs"
    "github.com/arturoguerra/goautoplex/internal/filebot"
    "github.com/arturoguerra/goautoplex/internal/nzbget"
    "github.com/arturoguerra/goautoplex/internal/deluge"
    "github.com/arturoguerra/goautoplex/internal/config"
)

type Api struct {
    Deluge *deluge.Deluge
    NzbGet *nzbget.NzbGet
}

func New(filebotcfg *config.FileBot, delugecfg *config.Deluge, nzbgetcfg *config.NzbGet) *Api {
    Filebot := filebot.New(filebotcfg)
    Deluge := deluge.New(delugecfg, Filebot)
    NzbGet := nzbget.New(nzbgetcfg, Filebot)

    return &Api{
        Deluge,
        NzbGet,
    }
}

func (a *Api) NzbGetHandler(c echo.Context) (err error) {
    d := new(structs.NzbGetPayload)
    if err = c.Bind(d); err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    if err = c.Validate(d); err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    a.NzbGet.Handle(d)

    return c.NoContent(http.StatusOK)
}

func (a *Api) DelugeHandler(c echo.Context) (err error) {
    d := new(structs.DelugePayload)
    if err = c.Bind(d); err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    if err = c.Validate(d); err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    go a.Deluge.Handle(d)

    return c.NoContent(http.StatusOK)
}
