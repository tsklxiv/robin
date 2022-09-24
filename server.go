// main.go - Contains the main stuff
package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// GET /
// displayFolder and notFound from helper_functions.go
func index(c echo.Context) error {
	path := ACCESSIBLE_DIR + c.QueryParam("path")
	log.Println("path = ", path)
	fi, err := os.Stat(path)
	if err != nil {
		return notFound(c, "File or directory not found")
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return displayFolder(c, path)
	case mode.IsRegular():
		return c.Inline(path, path)
	}

	return notFound(c, "Unknown error. Please try again.")
}

func serve(port int, dir string) {
	if dir != "" {
		ACCESSIBLE_DIR = dir
	}
  e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(spf(":%d", port)))
}
