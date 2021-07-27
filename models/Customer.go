package models

type Customer struct {
	CustomerID string `gorm:"primary_key;column:CustomerID"`
	CompanyName string `gorm:"column:CompanyName"`
	ContactName string `gorm:"column:ContactName"`
 	ContactTitle string `gorm:"column:ContactTitle"`
 	Address string `gorm:"column:Address"`
 	City string `gorm:"column:City"`
 	Region string `gorm:"column:Region"`
 	PostalCode string `gorm:"column:PostalCode"`
 	Country string `gorm:"column:Country"`
 	Phone string `gorm:"column:Phone"`
 	Fax string `gorm:"column:Fax"`
}
func (customer *Customer)TableName() string{
	return "Customers"
}