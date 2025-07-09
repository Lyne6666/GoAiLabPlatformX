// cmd/goailabplatformx/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"goailabplatformx/internal/goailabplatformx"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := goailabplatformx.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
