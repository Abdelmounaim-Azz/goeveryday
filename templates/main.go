package main

import (
	"bytes"
	"fmt"
	"text/template"
)

const letter = `Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
Best wishes,
Abdelmounaim`

func main() {
	people := []struct {
		Name     string
		Attended bool
	}{
		{
			Name:     "John",
			Attended: true,
		},
		{
			Name:     "Doe",
			Attended: false,
		},
	}
	t := template.Must(template.New("letter").Parse(letter))
	for _, person := range people {
		buffer := bytes.Buffer{}
		err := t.Execute(&buffer, person)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Letter for %s:\n\n%s\n\n", person.Name, buffer.String())
	}
}
