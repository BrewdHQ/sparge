package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func start(dir string, port int) {
	hasExtension := regexp.MustCompile(`\.[^\/]+$`)

	router := gin.Default()
	router.Use()
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.NoRoute(func(c *gin.Context) {
		if !hasExtension.MatchString(c.Request.URL.Path) {
			c.File("./public/index.html")
		}
	})
	router.Run(fmt.Sprintf(":%d", port))
}

func main() {
	app := cli.NewApp()
	app.Name = "hotspring"
	app.Usage = "A SPA (single-page application) server"
	app.Version = "1.0.0"

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
