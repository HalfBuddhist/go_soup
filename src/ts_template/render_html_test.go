package ts_template

import (
	"html/template"
	"net/http"
	"testing"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("ts_template/test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func TestTmpl(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()

}
