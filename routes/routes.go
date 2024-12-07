package routes

import (
	"net/http"

	"github.com/arthurazevedods/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
