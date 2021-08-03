package routes

import (
	"net/http"
	"projectGO/controllers"
)

func CarregaRota() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/New", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/deleta", controllers.Deleta)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
