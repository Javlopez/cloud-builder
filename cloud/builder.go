package cloud

import "fmt"

const Version string = "0.1"

type Builder struct {
	Version string
}

func New() *Builder {
	return &Builder{
		Version: Version,
	}
}

func (cb *Builder) Run(cmd string) {
	//Validate
	fmt.Printf("Running %s", cmd)
}
