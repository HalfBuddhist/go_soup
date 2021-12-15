package ts_template

import (
	"os"
	"testing"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func TestMain(t *testing.T) {
	p := Person{"longshuai",23}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
}
