package LibraryController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	libraryDbHelper "github.com/rubansundararaj/librarymanagement/dbhelper"
	librarymodel "github.com/rubansundararaj/librarymanagement/model"
)

func InsertOneBookRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var librarymodel librarymodel.LibraryModel
	_ = json.NewDecoder(r.Body).Decode(&librarymodel)

	status := libraryDbHelper.InsertOneBookToDb(librarymodel)

	if status {
		json.NewEncoder(w).Encode("success")
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("failed")
	}

}

func GetOneBookRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)

	status, books := libraryDbHelper.GetOneBookRecordFromDb(params["id"])
	if status {
		json.NewEncoder(w).Encode(books)
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("failed")
	}
}

func UpdateOneBookRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var librarymodel librarymodel.LibraryModel
	_ = json.NewDecoder(r.Body).Decode(&librarymodel)

	params := mux.Vars(r)

	status := libraryDbHelper.UpdateOneBookRecordInDb(params["id"], librarymodel)

	if status {
		json.NewEncoder(w).Encode("success")
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("failed")
	}

}

func DeleteOneBookRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)

	status := libraryDbHelper.DeleteOneBookRecordFromBb(params["id"])

	if status {
		json.NewEncoder(w).Encode("success")
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("failed")
	}
}

func GetAllBookRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	fmt.Println("In controller planning to get")
	status, books := libraryDbHelper.GetAllBookRecordsFromDb()
	if status {
		json.NewEncoder(w).Encode(books)
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("failed")
	}
}
