hotspring
-----------
<img src="https://raw.githubusercontent.com/BrewOpsHQ/hotspring/master/logo.jpg" align="right" height=175 width=175>
Create a SPA with hotspring. Hotspring is a Go based SPA (single page web application) server.

Hotspring is built with the following dependencies:
- [gin](https://github.com/gin-gonic/gin) 
- [gin-contrib/static](https://github.com/gin-contrib/static)
- [urfave/cli](https://github.com/urfave/cli)
- [glide](https://github.com/Masterminds/glide)

## Build

The following are instructions to build the hotspring binary. This project uses
[glide]() to manage dependencies.

```bash
$ cd $GOPATH
$ go get github.com/BrewOpsHQ/hotspring
$ cd hotspring
$ cd src/BrewOpsHQ/hotspring
$ glide install
$ go build
```

## Usage

Put `hostpring` in your path and run `hotspring help` for commands and flags. 

Commands:
```bash
$ hotspring help
NAME:
   hotspring - A SPA (single-page application) server

USAGE:
   hotspring.exe [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     start    Start the SPA server
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Start:
```bash
$ hotspring help start
NAME:
   hotspring.exe start - Start the SPA server

USAGE:
   hotspring.exe start [command options] [arguments...]

OPTIONS:
   --dir value, -d value   (default: "./public")
   --port value, -p value  (default: "8080")
```

Hotspring will serve assets from the `./public` directory by default. Be sure
to specifiy an alternate directory as needed.

## Version History

### 1.0.0

The first version

- Start server with a custom port
- Start server with a custom assets directory 