package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mnbao1975/microservices/product-images/files"
)

// Files is a handler for reading and writing files
type Files struct {
	store files.Storage
	log   *log.Logger
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{log: l, store: s}
}

// UploadREST implement the http.Handler interface
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Printf("POST - id: %s and filename: %s\n", id, fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters
	f.saveFile(id, fn, rw, r.Body)
}

// UploadMultipart something
func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Println("Bad request")
		http.Error(rw, "Expected multipart form", http.StatusBadRequest)
		return
	}

	_, idErr := strconv.Atoi(r.FormValue("id"))

	if idErr != nil {
		f.log.Println("Bad request")
		http.Error(rw, "Expected integer id", http.StatusBadRequest)
		return
	}

	mf, mh, errFile := r.FormFile("file")
	if errFile != nil {
		f.log.Println("Bad request")
		http.Error(rw, "Expected file", http.StatusBadRequest)
		return
	}
	f.saveFile(r.FormValue("id"), mh.Filename, rw, mf)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Printf("Save file for product id (%s) and path (%s)\n", id, path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Println("Unable to save file:", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "file was saved",
	})
}
