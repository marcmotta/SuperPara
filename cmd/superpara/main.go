// cmd/superpara/main.go
package main

import (
	"flag"
	"log"
	"os"

	"superpara/internal/superpara"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := superpara.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
