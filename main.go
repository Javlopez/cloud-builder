package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Javlopez/cloud-builder/args"
	"github.com/Javlopez/cloud-builder/spec/types"
)

func main() {
	_, err := args.NewArgReader(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("templates/basic/spec.yaml")
	if err != nil {
		log.Fatal(err)
	}

	spec, err := types.NewMainSpec(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", spec)
}
