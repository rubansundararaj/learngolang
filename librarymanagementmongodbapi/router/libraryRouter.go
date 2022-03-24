package libraryRouter

import (
	"github.com/gorilla/mux"
	LibraryController "github.com/rubansundararaj/librarymanagement/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/library/insertRecord", LibraryController.InsertOneBookRecord).Methods("POST")
	router.HandleFunc("/library/getall", LibraryController.GetAllBookRecords).Methods("GET")
	router.HandleFunc("/library/getone/{id}", LibraryController.GetOneBookRecord).Methods("GET")
	router.HandleFunc("/library/updateOne/{id}", LibraryController.UpdateOneBookRecord).Methods("PUT")
	router.HandleFunc("/library/deleteOne/{id}", LibraryController.DeleteOneBookRecord).Methods("DELETE")
	return router

}
