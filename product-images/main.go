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
	"github.com/mnbao1975/microservices/product-images/handlers"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Call the next handdler
		next.ServeHTTP(w, r)
	})
}

func ping() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "I am alive",
		})
	}
}

func main() {
	l := log.New(os.Stdout, "production-images: ", log.LstdFlags)

	hh := handlers.NewHello(l)
	sm := mux.NewRouter()
	sm.Use(commonMiddleware)

	sm.HandleFunc("/hello", hh.SayHello).Methods(http.MethodGet)
	// getR := sm.Methods(http.MethodGet).Subrouter()
	// getR.HandleFunc("/hello", hh.SayHello)

	// Ping endpoint in order to check server is avlive
	sm.HandleFunc("/ping", ping()).Methods("GET")

	sm.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "request not found",
		})
	})

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second, // the max time for connections using TCP Keep-Alive
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		l.Println("Server is running")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	fmt.Println("Receive terminate, graceful shutdown:", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}