package routing

import (
	"github.com/gorilla/mux"
	"goNorthwindApiDemo/controllers"
)

func Router(router *mux.Router){
	customerRoutes(router)
	productRoutes(router)
}

func customerRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/getall",controllers.CustomerController{}.GetCustomers).Methods("GET")
	router.HandleFunc("/api/customer/getbyid/{id}",controllers.CustomerController{}.GetCustomerById).Methods("GET")
	router.HandleFunc("/api/customer/create",controllers.CustomerController{}.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/customer/update",controllers.CustomerController{}.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/api/customer/delete/{id}",controllers.CustomerController{}.DeleteCustomerById).Methods("DELETE")
}

func productRoutes(router *mux.Router){
	router.HandleFunc("/api/product/getall",controllers.ProductController{}.GetProducts).Methods("GET")
	router.HandleFunc("/api/product/getbyid/{id}",controllers.ProductController{}.GetProductById).Methods("GET")
	router.HandleFunc("/api/product/create",controllers.ProductController{}.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/update",controllers.ProductController{}.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product/delete/{id}",controllers.ProductController{}.DeleteProduct).Methods("DELETE")
}