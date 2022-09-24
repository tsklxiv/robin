package main

import (
	"os"

	"github.com/mkideal/cli"
)

type Flags struct {
  cli.Helper
  Port int `cli:"p,port" usage:"Port number (default is 1323)" dft:"1323"`
  Directory string `cli:"d,dir" usage:"Directory to display (default is .)" dft:"./"`
}

func main() {
  os.Exit(cli.Run(new(Flags), func (ctx *cli.Context) error {
    argv := ctx.Argv().(*Flags)
    ctx.String("Port=%d, Directory=%s", argv.Port, argv.Directory)
    // Add a slash at the end of the directory path in case the user
    // didn't do so
    if argv.Directory[len(argv.Directory) - 1] != '/' {
      argv.Directory += "/"
    }
    serve(argv.Port, argv.Directory)
    return nil
  }))
}
