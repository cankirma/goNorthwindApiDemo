package data

import (
	"goNorthwindApiDemo/cfg"
	"goNorthwindApiDemo/models"
)
type CustomerModel struct {
}

// GetCustomer get all
func (customerModel CustomerModel) GetCustomer() ([]models.Customer, error) {
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
func (customerModel CustomerModel) GetCustomerById(id string)(models.Customer,error)  {
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
func (customerModel CustomerModel)  CreateCustomer(customer *models.Customer) error{
	db,err:= cfg.DbExport.OpenDB()
	if err!= nil {
	return err
	}else{
		db.Create(&customer)
		return nil

	}
}
//UpdateCustomer
func (customerModel CustomerModel) UpdateCustomer(customer *models.Customer)  error {
	db,err:= cfg.DbExport.OpenDB()
	if err!= nil {
		return err
	}else {
		db.Save(&customer)
		return nil
	}
}
func (customerModel CustomerModel) DeleteCustomer(customer models.Customer) error  {
	db,err := cfg.DbExport.OpenDB()
	if err!= nil{
		return err
	}else {
		db.Delete(customer)
		return nil
	}
}

