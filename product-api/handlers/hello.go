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

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//calling Println from the Logger of Hello
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	// Response to client
	fmt.Fprintf(rw, "Hello %s", d)
}
