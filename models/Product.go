package models

type Product struct {
Id int `gorm:"primary_key;column:ProductID"`
ProductName string `gorm:"column:ProductName"`
SupplierId int `gorm:"column:SupplierId"`
CategoryID int `gorm:"column:CategoryID"`
QuantityPerUnit string `gorm:"column:QuantityPerUnit"`
UnitPrice float64 `gorm:"column:UnitPrice"`
UnitsOnOrder int `gorm:"column:UnitsOnOrder"`
ReOrderLevel int `gorm:"column:ReOrderLevel"`
Discontinued bool `gorm:"column:Discontinued"`
}