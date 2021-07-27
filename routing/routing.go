package routing

import (
	"github.com/gorilla/mux"
	"goNorthwindApiDemo/controllers"
)

func Router(router *mux.Router){
	customerRoutes(router)
}

func customerRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/getall",controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/api/customer/getbyid/{id}",controllers.GetCustomerById).Methods("GET")
	router.HandleFunc("/api/customer/create",controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/customer/update",controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/api/customer/delete/{id}",controllers.DeleteCustomerById).Methods("DELETE")
}
