//Package handlers ...
//
//	Documentation for Product API
//		Schemes: http
//		BasePath: /
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//	swagger:meta
package handlers

import "github.com/mnbao1975/microservices/product-api/data"

//	A list of products return in the response
//	swagger:response productsResponse
type productsResponseWrapper struct {
	//	in: body
	Body []data.Product
}

//	A product returns in the response
//	swagger:response productResponse
type productResponseWrapper struct {
	//	in: body
	Body data.Product
}

//	An error returns in the response
//	swagger:response errorResponse
type errorResponseWrapper struct {
	//	in: body
	Body GenericError
}

//	No content is returned by this API endpoint
//	swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// in: body
	// required: true
	Body data.Product
}

//	swagger:parameters listOneProduct updateProduct
type productIDParamWrapper struct {
	//	The id of product given on the URI path
	//	in: path
	//	required: true
	ID int `json:"id"`
}
