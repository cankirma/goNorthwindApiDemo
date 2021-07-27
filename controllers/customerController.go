package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goNorthwindApiDemo/data"
	"goNorthwindApiDemo/models"
	"net/http"
)

//get
func GetCustomers(res http.ResponseWriter, req *http.Request) {
	var customerModel data.CustomerModel
	customers, err := customerModel.GetCustomer()
	if err != nil {
		responseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		responseWithJson(res, http.StatusOK, customers)
	}
}
func GetCustomerById(res http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	id := variables["id"]
	var customerModel = data.CustomerModel{}
	customers, err := customerModel.GetCustomerById(id)
	if err != nil {
		responseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		responseWithJson(res, http.StatusOK, customers)
	}
}
func CreateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		responseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		responseWithJson(res, http.StatusOK, customer)
	}
}
func UpdateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		responseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var customerModel data.CustomerModel
		err2 := customerModel.UpdateCustomer(&customer)
		if err2 != nil {
			responseWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(res, http.StatusOK, customer)
		}
	}

}
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]
	var customerModel data.CustomerModel
	customers, err := customerModel.GetCustomerById(id)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	} else {
		responseWithJson(w, http.StatusOK, customers)
	}
}
func responseWithError(w http.ResponseWriter, statusCode int, msg string) {
	responseWithJson(w, statusCode, map[string]string{"error": msg})
}
func responseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "Application-Json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
