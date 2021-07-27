package data

import (
	"goNorthwindApiDemo/cfg"
	"goNorthwindApiDemo/models"
)

type ProductData struct {

}

func (productData ProductData) GetProducts() ([]models.Product,error){
	db,err := cfg.DbExport.OpenDB()
	if err!= nil{
		return nil, err
	}else {
		var products []models.Product
		db.Find(&products)
		return products, err
	}
}

func (productData ProductData) GetProductById(id int) (models.Product, error) {
db,err := cfg.DbExport.OpenDB()
	if err!= nil {
		return models.Product{},  err
	}else {
		var product models.Product
		db.Where("ProductID like ?",id).Find(&product)
		return product, nil
	}
}
func (productData ProductData) CreateProduct(product models.Product) error {
	db, err := cfg.DbExport.OpenDB()
	if err != nil{
		return err
	}else {
		db.Create(&product)
		return nil
	}
}
func(productData ProductData) UpdateProduct(product models.Product) error  {
	db, err := cfg.DbExport.OpenDB()
	if err!= nil {
		return err
	}else {
		db.Save(&product)
		return nil
	}
}
func (productData ProductData) DeleteProduct(product models.Product) error{
	db,err := cfg.DbExport.OpenDB()
	if err!= nil{
		return err
	}else {
		db.Delete(product)
	return nil
	}
}

