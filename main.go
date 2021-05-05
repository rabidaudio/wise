package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "wise",
		Description: "Smarter version control than git",
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
