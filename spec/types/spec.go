package types

import (
	"errors"

	"gopkg.in/yaml.v2"
)

//MainSpec struct is the base of the spec files
type MainSpec struct {
	//Version of the spec
	Version string `yaml:"Version"`
	//Engine to be processs (aws/gpc/Azure/Alibaba)
	Engine string `yaml:"Engine"`
	//Class what kind of operation needs to perform
	Class string `yaml:"Class"`
	//Name of the task
	Name string `yaml:"Name"`
	//Spec TODO
	Spec string `yaml:"Spec"`
	//Output TODO
	Output string `yaml:"Output"`
}

//NewMainSpec TODO
func NewMainSpec(data []byte) (MainSpec, error) {
	var spec MainSpec
	err := yaml.Unmarshal(data, &spec)
	if err != nil {
		return MainSpec{}, err
	}

	err = Validate(spec)
	if err != nil {
		return MainSpec{}, err
	}

	//must to run a validator
	return spec, nil
}

const (
	ErrInvalidVersion        = "The version is not valid"
	ErrInvalidCloudProvider  = "The cloud provider is not supported yet"
	ErrInvalidClassOperation = "The class is invalid or is empty"
	ErrInvalidNameTask       = "The name of the spec task cannot be empty"
)

//Validator...
func Validate(spec MainSpec) error {
	if spec.Version != "0.0.1" {
		return errors.New(ErrInvalidVersion)
	}
	if spec.Engine != "aws" {
		return errors.New(ErrInvalidCloudProvider)
	}
	if spec.Class != "ec2" {
		return errors.New(ErrInvalidClassOperation)
	}
	if spec.Name == "" {
		return errors.New(ErrInvalidNameTask)
	}
	return nil
}
