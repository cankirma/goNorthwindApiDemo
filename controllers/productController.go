package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goNorthwindApiDemo/data"
	"goNorthwindApiDemo/models"
	"goNorthwindApiDemo/utils"
	"net/http"
	"strconv"
)

type ProductController struct {
	data.ProductData
}


func(ProductController)GetProducts(res http.ResponseWriter,req *http.Request){
	var productData data.ProductData
	products,err := productData.GetProducts()
	if err !=nil {
		utils.ResponseWithError(res,http.StatusBadRequest,err.Error())
	}else {
		utils.ResponseWithJson(res,http.StatusOK,products)
	}
}
func (ProductController)GetProductById(res http.ResponseWriter,req *http.Request){
	variables := mux.Vars(req)
	id := variables["id"]
	productId , _ := strconv.Atoi(id)
	var productData = data.ProductData{}
	customers ,err := productData.GetProductById(productId)
	if err!= nil {
		utils.ResponseWithError(res,http.StatusBadRequest,err.Error())
	}else {
		utils.ResponseWithJson(res,http.StatusOK,customers)
	}
}
func (ProductController) CreateProduct(res http.ResponseWriter,req *http.Request)  {
	var product models.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		utils.ResponseWithError(res,http.StatusBadRequest,err.Error())
	}else {
		var productData data.ProductData
		err2 := productData.CreateProduct(product)
		if err2 != nil {
			utils.ResponseWithError(res,http.StatusBadRequest,err2.Error())
		}else {
			utils.ResponseWithJson(res,http.StatusOK,product)
		}
	}
}

func ( ProductController) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	var product models.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		utils.ResponseWithError(writer,http.StatusBadRequest,err.Error())
	}else {
		var productData data.ProductData
		err2 := productData.UpdateProduct(product)
		if err2!= nil {
			utils.ResponseWithError(writer,http.StatusBadRequest,err2.Error())
		}else{
			productData.UpdateProduct(product)
		}
	}
}

func (c ProductController) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
variables := mux.Vars(request)
id := variables["id"]
productId,_ := strconv.Atoi(id)
var productData data.ProductData
product ,err := productData.GetProductById(productId)
	if err!=nil {
		utils.ResponseWithError(writer,http.StatusOK,err.Error())
	}else {
		err2:=  productData.DeleteProduct(product)
		if err2 != nil{
			utils.ResponseWithError(writer,http.StatusBadRequest,err.Error())
		}else{
			utils.ResponseWithJson(writer,http.StatusOK,product)
		}
	}
}


