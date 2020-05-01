package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIvalidYamlSpec(t *testing.T) {
	for _, tc := range []struct {
		Title string
		Yaml  string
		Error string
	}{
		{Title: "Should avoid any file not valid", Yaml: "Something", Error: "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Something` into types.MainSpec"},
		{Title: "The version should be any valid version", Yaml: "Version: 0.0.0", Error: "The version is not valid"},
		{Title: "The engine must be supported", Yaml: "Version: 0.0.1\nEngine: myProvider", Error: "The cloud provider is not supported yet"},
		{Title: "The class must be supported", Yaml: "Version: 0.0.1\nEngine: aws\nClass: anything", Error: "The class is invalid or is empty"},
		{Title: "Must be specified any name for the spec task", Yaml: "Version: 0.0.1\nEngine: aws\nClass: ec2\nName:  ", Error: "The name of the spec task cannot be empty"},
	} {
		t.Run(tc.Title, func(t *testing.T) {
			yaml := []byte(tc.Yaml)
			_, err := NewMainSpec(yaml)
			assert.EqualError(t, err, tc.Error)
		})

	}
}

func TestValidSpecYamlFile(t *testing.T) {

	t.Run("The operation can perform once the spec is valid", func(t *testing.T) {
		yamlAsString := "Version: 0.0.1\nEngine: aws\nClass: ec2\nName: New_Machine"
		validYaml := []byte(yamlAsString)
		spec, err := NewMainSpec(validYaml)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "0.0.1", spec.Version)
		assert.Equal(t, "aws", spec.Engine)
		assert.Equal(t, "ec2", spec.Class)
		assert.Equal(t, "New_Machine", spec.Name)
	})
}
