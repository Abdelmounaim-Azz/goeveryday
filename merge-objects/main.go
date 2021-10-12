package main

import (
	"fmt"

	"github.com/imdario/mergo"
	yaml "gopkg.in/yaml.v2"
)

type Spec struct {
	// Name name of the function
	Name string `yaml:"name"`
	// Image docker image name of the function
	Image       string            `yaml:"image"`
	Environment map[string]string `yaml:"environment,omitempty"`
	// Limits for the function
	Limits *FunctionResources `yaml:"limits,omitempty"`
}

// FunctionResources Memory and CPU
type FunctionResources struct {
	Memory string `yaml:"memory"`
	CPU    string `yaml:"cpu"`
}

func main() {
	base := Spec{
		Image: "docker.io/functions/figlet:latest",
		Name:  "figlet",
	}
	production := Spec{
		Environment: map[string]string{"stage": "production"},
		Limits:      &FunctionResources{Memory: "1Gi", CPU: "100Mi"},
	}
	overrides := []Spec{
		base,
		production,
	}
	merged := Spec{}
	for _, override := range overrides {
		err := mergo.Merge(&merged, override, mergo.WithOverride)
		if err != nil {
			panic(err)
		}
	}
	bytesOut, err := yaml.Marshal(merged)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Merged content:\n\n%s\n", string(bytesOut))
}
