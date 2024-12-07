package main

import (
	"html/template"
	"net/http"

	"github.com/arthurazevedods/models"
	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
