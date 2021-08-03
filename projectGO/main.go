package main

import (
	"net/http"
	"projectGO/routes"
)

func main() {
	routes.CarregaRota()
	http.ListenAndServe(":8000", nil)

}
