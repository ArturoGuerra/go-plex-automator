package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/arturoguerra/goautoplex/internal/api"
    "github.com/arturoguerra/goautoplex/internal/config"
    "github.com/arturoguerra/goautoplex/internal/datavalidator"
    "os"
)

var (
    filebotcfg *config.FileBot
    delugecfg *config.Deluge
    nzbgetcfg *config.NzbGet
)

func init () {
    filebotcfg = &config.FileBot{
        os.Getenv("FILEBOT_DESTDIR"),
        os.Getenv("FILEBOT_AMCLOGS"),
        os.Getenv("FILEBOT_AMCEXCLUDE"),
    }

    delugecfg = &config.Deluge{
        os.Getenv("DELUGE_BASEDIR"),
        os.Getenv("DELUGE_APIURL"),
        os.Getenv("DELUGE_APIKEY"),
    }

    nzbgetcfg = &config.NzbGet{
        os.Getenv("NZBGET_BASEDIR"),
    }
}

func main () {
    apiHandler := api.New(filebotcfg, delugecfg, nzbgetcfg)

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Validator = datavalidator.New()

    e.POST("/nzbget", apiHandler.NzbGetHandler)
    e.POST("/deluge", apiHandler.DelugeHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    e.Logger.Fatal(e.Start(":" + port))
}
