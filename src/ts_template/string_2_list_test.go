package ts_template

import (
	"os"
	"text/template"
)

func TestMain2() {
	tmpl, _ := template.New("test").Parse(`{{ (split "." .)._0 }}`)
	_ = tmpl.Execute(os.Stdout, "hello.world")
}
