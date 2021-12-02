package main

import (
	"net/http"

	"github.com/flaviofilipe/loja_go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
