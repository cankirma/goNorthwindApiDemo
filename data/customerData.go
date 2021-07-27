package data

import (
	"goNorthwindApiDemo/cfg"
	"goNorthwindApiDemo/models"
)
type CustomerData struct {
}

// GetCustomer get all
func (customerData CustomerData) GetCustomer() ([]models.Customer, error) {
	db ,err := cfg.DbExport.OpenDB()
	if err != nil {
		return nil,err
	}else {
		var customers []models.Customer
		db.Find(&customers)
		return customers,nil
	}
}

// GetCustomerById get by id get/{id}
func (customerData CustomerData) GetCustomerById(id string)(models.Customer,error)  {
	db,err := cfg.DbExport.OpenDB()
	if err!= nil{
		return models.Customer{}, err
	}else {
		var customer models.Customer
		db.Where("CustomerID like ?",id).Find(&customer)
		return customer ,nil
	}
}

// CreateCustomer post
func (customerData CustomerData)  CreateCustomer(customer *models.Customer) error{
	db,err:= cfg.DbExport.OpenDB()
	if err!= nil {
	return err
	}else{
		db.Create(&customer)
		return nil

	}
}
//UpdateCustomer
func (customerData CustomerData) UpdateCustomer(customer *models.Customer)  error {
	db,err:= cfg.DbExport.OpenDB()
	if err!= nil {
		return err
	}else {
		db.Save(&customer)
		return nil
	}
}
func (customerData CustomerData) DeleteCustomer(customer models.Customer) error  {
	db,err := cfg.DbExport.OpenDB()
	if err!= nil{
		return err
	}else {
		db.Delete(customer)
		return nil
	}
}

