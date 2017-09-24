package main

import (
	"net/http"

	"./model"
	"./providers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func getRouter(db *gorm.DB, p providers.StorageProvider) *mux.Router {
	r := mux.NewRouter()

	f := r.PathPrefix("/files").Subrouter()
	f.HandleFunc("/{name}", uploadFile(db, p)).Methods("POST")
	f.HandleFunc("/{name}", deleteFile(db, p)).Methods("DELETE")

	return r
}

func deleteFile(db *gorm.DB, p providers.StorageProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		file := model.File{}
		err := db.Where("name = ?", name).First(&file).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err := p.DeleteFile(name); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = db.Model(&model.File{}).Delete(&model.File{}, name).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusPartialContent)
		return
	}
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
