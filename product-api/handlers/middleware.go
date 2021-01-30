package handlers

import (
	"context"
	"net/http"

	"github.com/mnbao1975/microservices/product-api/data"
)

// MiddlewareValidateProduct is used to validate product data
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		//err := prod.FromJSON(r.Body)
		err := data.FromJSON(&prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] decentializing product")
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(GenericError{Message: err.Error()}, rw)
			return
		}

		// Validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(GenericError{Message: err.Error()}, rw)

			return
		}

		// Copy the valid product(prod) to the KeyProduct and add it to the context
		// so that the next handdler can access it
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handdler
		next.ServeHTTP(rw, r)
	})
}
