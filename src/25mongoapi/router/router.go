package router

import (
	"github.com/gorilla/mux"
	"github.com/kinjalrk2k/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/movie", controller.Create).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.DeleteAll).Methods("DELETE")

	return router
}
