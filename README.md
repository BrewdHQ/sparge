hotspring
-----------
<img src="https://raw.githubusercontent.com/BrewOpsHQ/hotspring/master/hotspring-logo.png" align="right" height=175 width=175>

Create a SPA with hotspring. Hotspring is a Go based SPA (single page web application) 
server.

The purpose of hotspring is to allow single page frameworks like Aurelia, 
Angular, and React to use "pretty urls" (urls without the #) by serving the index.html
file from the assets directory whenever a static file is not found. Hotspring will,
however, return 404 if an extension is used in the path, such as `favicon.png`, and 
that file is not found.

1. Look for the path in the assets directory and return that static asset if it exists.
2. If the path contains no extension (e.g., `.html`, `.png`, `.jpg`), then return `index.html`
from the assets directory.
3. Otherwise, return a 404.

## Dependencies

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
   hotspring [global options] command [command options] [arguments...]

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
   hotspring start - Start the SPA server

USAGE:
   hotspring start [command options] [arguments...]

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