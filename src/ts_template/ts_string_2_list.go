package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, _ := template.New("test").Parse(`{{ (split "." .)._0 }}`)
	_ = tmpl.Execute(os.Stdout, "hello.world")
}
