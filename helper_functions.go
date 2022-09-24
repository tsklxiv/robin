// helper_functions.go - Helper functions
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/olekukonko/tablewriter"
)

func displayFolder(c echo.Context, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"TYPE", "NAME", "SIZE", "LAST MODIFIED"})
	for _, file := range files {
		lastModified := file.ModTime().Format("Mon Jan _2 15:04:05 2006")
		if file.IsDir() {
			table.Append([]string{"DIR", file.Name(), "", lastModified})
		} else {
			table.Append([]string{"FILE", file.Name(), ByteCountSI(file.Size()), lastModified})
		}
	}
	table.Render()

	return c.String(OK, tableString.String() + HELP)
}

func notFound(c echo.Context, msg string) error {
	return c.String(NOT_FOUND, msg)
}

// Formatting byte size in human-readable format (SI/decimal)
// From https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func ByteCountSI(b int64) string {
    const unit = 1000
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB",
        float64(b)/float64(div), "kMGTPE"[exp])
}
