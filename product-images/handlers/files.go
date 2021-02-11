package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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

func (f *Files) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Printf("POST - id: %s and filename: %s\n", id, fn)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "file was saved",
	})

}
