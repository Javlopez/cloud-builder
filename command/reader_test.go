package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	t.Run("If any argument is passed should return an error", func(t *testing.T) {
		_, err := NewArgReader([]string{})
		assert.Error(t, err)
	})

	t.Run("Only one command will be available", func(t *testing.T) {
		_, err := NewArgReader([]string{"test", "apply"})
		assert.Error(t, err)
	})
}

/*
ArgsReader := NewArgReader
NewArgReader.GetCommand()

x:= NewArgReader.GetCmdValue("demo")// x = 1
argsWithoutProg := os.Args[1:]

cloud-builder validate infra.yaml
cloud-builder apply infra.yaml
cloud-builder version*/
