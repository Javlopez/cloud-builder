package main

import (
	"log"
	"os"

	cb "github.com/Javlopez/cloud-builder/cloud"
	"github.com/Javlopez/cloud-builder/command"
)

func main() {
	args, err := command.NewArgReader(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	app := cb.New()
	app.Run(args.GetCommand())
}
