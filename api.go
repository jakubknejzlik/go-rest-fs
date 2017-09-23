package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakubknejzlik/go-rest-fs/providers"
	"github.com/jinzhu/gorm"
)

func getRouter(db *gorm.DB, p providers.StorageProvider) *mux.Router {
	r := mux.NewRouter()

	f := r.PathPrefix("/files").Subrouter()
	f.HandleFunc("/{name}", uploadFile(db, p)).Methods("POST")

	return r
}

func uploadFile(db *gorm.DB, p providers.StorageProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		data := r.Body
		if err := p.UploadFile(data, name); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
