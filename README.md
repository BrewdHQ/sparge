sparge
------
Sparge is a Go based SPA (single page web application) 
server.

The purpose of sparge is to allow single page frameworks like Aurelia, 
Angular, and React to use "pretty urls" (urls without the #) by serving the index.html
file from the assets directory whenever a static file is not found. Sparge will,
however, return 404 if an extension is used in the path, such as `favicon.png`, and 
that file is not found.

1. Look for the path in the assets directory and return that static asset if it exists.
2. If the path contains no extension (e.g., `.html`, `.png`, `.jpg`), then return `index.html`
from the assets directory.
3. Otherwise, return a 404.

## Dependencies

Sparge is built with the following dependencies:
- [labstack/echo](https://github.com/labstack/echo)
- [urfave/cli](https://github.com/urfave/cli)
- [glide](https://github.com/Masterminds/glide)

## Build

The following are instructions to build the sparge binary. This project uses
[glide]() to manage dependencies.

```bash
$ cd $GOPATH
$ go get github.com/BrewdHQ/sparge
$ cd sparge
$ cd src/BrewdHQ/sparge
$ glide install
$ go build
```

## Usage

Put `sparge` in your path and run `sparge help` for commands and flags. 

Commands:
```bash
$ sparge help
NAME:
   sparge - A SPA (single-page application) server

USAGE:
   sparge [global options] command [command options] [arguments...]

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
$ sparge help start
NAME:
   sparge start - Start the SPA server

USAGE:
   sparge start [command options] [arguments...]

OPTIONS:
   --dir value, -d value   (default: "./public")
   --port value, -p value  (default: "8080")
```

Sparge will serve assets from the `./public` directory by default. Be sure
to specifiy an alternate directory as needed.

## Version History

### 1.0.0

The first version

- Start server with a custom port
- Start server with a custom assets directory 