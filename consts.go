// consts.go - Constants

package main

import (
  "fmt"
  "net/http"
)

const OK int = http.StatusOK
const NOT_FOUND int = http.StatusNotFound
const HELP = `
HELP
===================================
Robin is a simple and terminal-friendly file server written in Go.

To read/download files available from the server, go to /?path=<path>,
with <path> being the file you want to read.

For example, to read the file abc.txt in directory def, <path> should
be:

  /?path=def/abc.txt
`
var ACCESSIBLE_DIR = "./"
var spf = fmt.Sprintf
var sp = fmt.Sprint
