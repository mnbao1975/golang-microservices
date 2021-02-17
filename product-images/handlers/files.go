package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

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

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Printf("POST - id: %s and filename: %s\n", id, fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters
	f.saveFile(id, fn, rw, r)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Printf("Save file for product id (%s) and path (%s)\n", id, path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
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
