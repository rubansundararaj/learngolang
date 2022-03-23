package router

import (
	netFlixController "mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", netFlixController.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", netFlixController.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}", netFlixController.MarkAsWatched).Methods("PUT")

	router.HandleFunc("/api/movie/{id}", netFlixController.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", netFlixController.DeleteAllMovies).Methods("DELETE")

	return router
}
