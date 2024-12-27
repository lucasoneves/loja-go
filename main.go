package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	fmt.Println("Bem-vindo a loja Go")

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}

	err := htmlTemplates.ExecuteTemplate(w, "index.html", produtos)
	if err != nil {
		http.Error(w, "Erro ao renderizar o template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
