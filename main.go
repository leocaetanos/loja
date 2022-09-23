package main

import (
	"net/http"

	"github.com/leocaetanos/loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
