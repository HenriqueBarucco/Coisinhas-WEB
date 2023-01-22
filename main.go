package main

import (
	"net/http"

	"github.com/henriquebarucco/Golang-WEB/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
