package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is
type Hello struct {
	l *log.Logger
}

// NewHello is pattern of dependency injection, a logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// SayHello print out hello
func (h *Hello) SayHello(w http.ResponseWriter, r *http.Request) {
	//calling Println from the Logger of Hello
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	// Response to client
	fmt.Fprintf(w, "Hello %s", d)
}
