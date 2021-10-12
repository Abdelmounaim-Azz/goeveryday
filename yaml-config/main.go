package main

import (
	"fmt"
	"io/ioutil"
	"os"

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
	// bytesOut, err := ioutil.ReadFile("config.yaml")
	// if err != nil {
	// 	panic(err)
	// }
	// spec := Spec{}
	// if err := yaml.Unmarshal(bytesOut, &spec); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Function name: %s\tImage: %s\tEnvs: %d\n", spec.Name, spec.Image, len(spec.Environment))
	spec := Spec{
		Image: "docker.io/functions/figlet:latest",
		Name:  "figlet",
	}
	bytesOut, err := yaml.Marshal(spec)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("figlet.yaml", bytesOut, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote: figlet.yaml.. OK.\n")
}
