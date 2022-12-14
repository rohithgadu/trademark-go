package routes

import (
	"github.com/gorilla/mux"
	"go-demo/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/trademark", controllers.CreateRecord).Methods("POST")
	router.HandleFunc("/trademark/{recordId}", controllers.GetRecordById).Methods("GET")

}
