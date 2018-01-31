package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/urfave/cli"
)

const (
	banner = `
 _____ _____ _____             
|   __|  _  |  _  |___ ___ ___ 
|__   |   __|     |  _| . | -_|
|_____|__|  |__|__|_| |_  |___| %s
                      |___|
Single Page Application server  ________
_______________________________/__\__\__\
                               \__/__/__/

`
)

// version set by LDFLAGS
var version string

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   root,
		Index:  "index.html",
		HTML5:  true,
		Browse: false,
	}))

	fmt.Printf(banner, "v"+version)
	fmt.Printf("» http server started on port %d\n", port)
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func main() {
	app := cli.NewApp()
	app.Name = "sparge"
	app.Usage = "A SPA (single-page application) server"
	app.Version = version

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the SPA server",
			Action: func(c *cli.Context) error {
				start(c.String("dir"), c.Int("port"))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir, d",
					Value: "./public",
				},
				cli.StringFlag{
					Name:  "port, p",
					Value: "8080",
				},
			},
		},
	}

	// start()
	app.Run(os.Args)
}
