package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Bem-vindo a loja Go")

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	htmlTemplates.Execute(w, "index")
}
