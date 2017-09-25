package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

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
			return
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
		_, headerMultipart, err := r.FormFile("file")
		if err == nil {
			err, status := uploadFileMultipart(db, p, data, headerMultipart, name)
			if err != nil {
				http.Error(w, err.Error(), status)
			}
			return
		}
		headerData := r.Header
		err, status := uploadFileData(db, p, data, headerData, name)
		if err != nil {
			http.Error(w, err.Error(), status)
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}

func uploadFileData(db *gorm.DB, p providers.StorageProvider, r io.ReadCloser, h http.Header, name string) (error, int) {
	size, _ := strconv.Atoi(h.Get("Content-Length"))

	if err := p.UploadFile(r, name); err != nil {
		return err, http.StatusBadRequest
	}

	if err := model.CreateFileInDB(db, name, uint(size)); err != nil {
		return err, http.StatusBadRequest
	}
	return nil, http.StatusOK
}

func uploadFileMultipart(db *gorm.DB, p providers.StorageProvider, r io.ReadCloser, h *multipart.FileHeader, name string) (error, int) {
	var size int64
	mfile, _ := h.Open()
	switch t := mfile.(type) {
	case *os.File:
		fi, _ := t.Stat()
		size = fi.Size()
	default:
		size, _ = mfile.Seek(0, 0)
	}

	if err := p.UploadFile(r, name); err != nil {
		return err, http.StatusBadRequest
	}

	if err := model.CreateFileInDB(db, name, uint(size)); err != nil {
		return err, http.StatusBadRequest
	}
	return nil, http.StatusOK
}
