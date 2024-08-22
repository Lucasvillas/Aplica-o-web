package main

import (
	"go-loja-master/routes"
	"net/http"
)

func main() {
	routes.Rotas()
	http.ListenAndServe(":8080", nil)
}
