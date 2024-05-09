package repository

import "github.com/Adekabang/eniqilo-store/model"

type ProductRepositoryInterface interface {
	AddProduct(model.AddUpdateProduct) bool
	GetProduct(model.ParamsGetProduct) bool
	UpdateProduct(model.AddUpdateProduct) bool
	DeleteProduct(model.ProductUri) bool
}
