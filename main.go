package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/labstack/echo/v4"
)

var (
	app        = kingpin.New("App", "Simple app")
	argAppName = app.Flag("name", "Application Name").Short('n').Required().String()
	argPort    = app.Flag("port", "Server Port").Short('p').Default("9000").Int()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	appName := *argAppName

	port := fmt.Sprintf(":%d", *argPort)

	fmt.Printf("Starting %s at %s", appName, port)

	e := echo.New()
	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(port))
}
