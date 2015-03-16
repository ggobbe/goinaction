package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/ggobbe/goinaction/chapter2/sample/matchers"
	"github.com/ggobbe/goinaction/chapter2/sample/search"
)

var start time.Time

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
	start = time.Now()
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")

	fmt.Printf("Time elapsed: %s\n", time.Now().Sub(start))
}
