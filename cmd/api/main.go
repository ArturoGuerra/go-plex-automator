package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/arturoguerra/goautoplex/internal/api"
    "os"
)

func main () {
    e := echo.New()
    e.Use(Middleware.Logger())
    e.Use(Middleware.Recover())

    e.POST("/nzbget", api.NzbGetHandler)
    e.POST("/deluge", api.DelugeHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    e.Logger.Fatal(e.Start(":" + port))
}
