package router

import (
	"fmt"
	"userinterface/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Routers initialized")
	r := mux.NewRouter()
	r.HandleFunc("/upload", controller.Upload).Methods("POST")
	r.HandleFunc("/download/{name}", controller.Download).Methods("GET")
	return r
}
