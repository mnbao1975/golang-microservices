package handlers

import (
	"log"
	"net/http"
)

// GoodBye ...
type GoodBye struct {
	l *log.Logger
}

// NewGoodBye ...
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye")
	rw.Write([]byte("Byeeee"))
}
