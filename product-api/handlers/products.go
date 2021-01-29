package handlers

import (
	"log"
)

// KeyProduct is the key used for the Product in the context
type KeyProduct struct{}

// Products defines ...
type Products struct {
	l *log.Logger
}

// NewProducts is pattern of dependency injection, a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
