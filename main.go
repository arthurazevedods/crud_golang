package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) // retornar todos os templates html

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	//slice de produtos
	produtos := []Produto{
		{Nome: "Camisa", Descricao: "Básica", Preco: 29.99, Quantidade: 10},
		{Nome: "Garrafa térmica", Descricao: "300ml", Preco: 59.99, Quantidade: 2},
		{"Copo", "Corrida", 224.99, 5},
		{"Fone", "Headphone", 99.99, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
