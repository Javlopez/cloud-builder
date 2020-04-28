package command

import (
	"errors"
)

const ErrArgumentIsNotDefined = "The argument must be definined"
const ErrArgumentIsNotValid = "The argument is not valid"

var cmdAvailable = []string{"validate", "plan", "apply"}

//ArgReader struct
type ArgReader struct {
	Command string
}

//GetCommand method
func (args *ArgReader) GetCommand() string {
	return args.Command
}

func validateCommand(arg string) bool {
	for _, cmd := range cmdAvailable {
		if cmd == arg {
			return true
		}
	}
	return false
}

//NewArgReader constructor
func NewArgReader(args []string) (*ArgReader, error) {
	if len(args) == 0 {
		return &ArgReader{}, errors.New(ErrArgumentIsNotDefined)
	}
	command := args[0]

	if !validateCommand(command) {
		return &ArgReader{}, errors.New(ErrArgumentIsNotValid)
	}
	return &ArgReader{
		Command: command,
	}, nil
}
