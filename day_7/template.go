package main

import (
	"fmt"
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t, err := template.New("todos").
		Parse(`You have a task named "{{ .Name}}" with description: "{{ .Description}}"`)
	if err != nil {
		panic(err)
	}

	if err = t.Execute(os.Stdout, td); err != nil {
		panic(err)
	}

	fmt.Println()
	counter := 0
	// -- new counter
	const templateWithCounter = `Start Counter is: {{getCounter}}.{{incrementCounter}}You have a task named "{{ .Name}}" with description: "{{ .Description}}. Counter is: {{getCounter}}"`
	t2 := template.Must(template.New("todos2").Funcs(template.FuncMap{
		"incrementCounter": func() string {
			counter++
			return ""
		},
		"getCounter": func() int {
			return counter
		},
	}).Parse(templateWithCounter))
	if err = t2.Execute(os.Stdout, td); err != nil {
		panic(err)
	}
}
