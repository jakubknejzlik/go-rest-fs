package main

import (
	"net/http"

	"./model"
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
		if r.ContentLength < 1 {
			http.Error(w, "ContentLength is not set!", http.StatusBadRequest)
		}

		if err := p.UploadFile(data, name); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := model.CreateFileInDB(db, name, uint(r.ContentLength)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
