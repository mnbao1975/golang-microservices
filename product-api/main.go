package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mnbao1975/microservices/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "production-api: ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", ph.GetProducts)
	getR.HandleFunc("/{id:[0-9]+}", ph.GetOneProduct)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/", ph.AddProduct)
	postR.Use(ph.MiddlewareValidateProduct)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putR.Use(ph.MiddlewareValidateProduct)

	sm.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "request not found",
		})
	})

	s := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second, // the max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second}

	go func() {
		fmt.Println("Server is running")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Receive terminate, graceful shutdown:", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
