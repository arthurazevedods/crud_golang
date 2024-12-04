package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"

	_ "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) // retornar todos os templates html

func main() {
	db := conectaComBanco()
	defer db.Close()
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

func conectaComBanco() *sql.DB {
	// Obtém as variáveis de ambiente para conexão
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PWD")
	dbname := "golang_loja" // Nome fixo do banco (pode ser parametrizado)
	host := "localhost"     // Host fixo (também pode ser parametrizado)
	sslmode := "disable"    // Normalmente "disable" para desenvolvimento local

	// Cria a string de conexão no formato correto
	conexao := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
