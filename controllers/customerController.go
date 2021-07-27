package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goNorthwindApiDemo/data"
	"goNorthwindApiDemo/models"
	"goNorthwindApiDemo/utils"
	"net/http"
)
type CustomerController struct {

}
//get
func (customerController CustomerController) GetCustomers(res http.ResponseWriter, req *http.Request) {
	var customerModel data.CustomerData
	customers, err := customerModel.GetCustomer()
	if err != nil {
		utils.ResponseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseWithJson(res, http.StatusOK, customers)
	}
}
func  (customerController CustomerController) GetCustomerById(res http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	id := variables["id"]
	var customerData = data.CustomerData{}
	customers, err := customerData.GetCustomerById(id)
	if err != nil {
		utils.ResponseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseWithJson(res, http.StatusOK, customers)
	}
}
func  (customerController CustomerController) CreateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		utils.ResponseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseWithJson(res, http.StatusOK, customer)
	}
}
func  (customerController CustomerController) UpdateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		utils.ResponseWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var customerModel data.CustomerData
		err2 := customerModel.UpdateCustomer(&customer)
		if err2 != nil {
			utils.ResponseWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			utils.ResponseWithJson(res, http.StatusOK, customer)
		}
	}

}
func(customerController CustomerController) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]
	var customerData data.CustomerData
	customers, err := customerData.GetCustomerById(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseWithJson(w, http.StatusOK, customers)
	}
}

