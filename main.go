package main

import (
	"fmt"
	"os"
	"time"

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
	timeout = "620s"
)

// version set by LDFLAGS
var version string

func start(root string, port int, redirectHttps bool, logFormat string) {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat,
	}))

	if redirectHttps {
		e.Pre(middleware.HTTPSRedirect())
	}

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   root,
		Index:  "index.html",
		HTML5:  true,
		Browse: false,
	}))
	timeOut, _ := time.ParseDuration("620s")
	e.Server.IdleTimeout = timeOut

	fmt.Printf(banner, "v"+version)
	fmt.Printf("» http server started on port %d\n", port)

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
				start(c.String("dir"), c.Int("port"), c.Bool("https-redirect"), c.String("log-format")+"\n")
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "dir, d",
					Value:  "./public",
					Usage:  "Directory from which to serve files.",
					EnvVar: "SPARGE_DIR",
				},
				cli.StringFlag{
					Name:   "port, p",
					Value:  "8080",
					Usage:  "Server port.",
					EnvVar: "SPARGE_PORT",
				},
				cli.BoolFlag{
					Name:   "https-redirect",
					Usage:  "Use to force http to redirect to https",
					EnvVar: "SPARGE_HTTPS_REDIRECT",
				},
				cli.StringFlag{
					Name: "log-format",
					Value: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
						`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
						`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
						`"bytes_out":${bytes_out}}`,
					Usage:  "Specify the log format",
					EnvVar: "SPARGE_LOG_FORMAT",
				},
			},
		},
	}

	// start()
	app.Run(os.Args)
}
