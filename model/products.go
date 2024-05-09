package model

type AddUpdateProduct struct {
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Category    string `json:"category"`
	ImageUrl    string `json:"imageUrl"`
	Notes       string `json:"notes"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Location    string `json:"location"`
	IsAvailable bool   `json:"isAvailable"`
}

type ParamsGetProduct struct {
	Id          string `json:"id"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	Name        string `json:"Name"`
	IsAvailable bool   `json:"isAvailable"`
	Category    string `json:"category"`
	Sku         string `json:"sku"`
	Price       string `json:"price"`
	InStock     bool   `json:"inStock"`
	CreatedAt   string `json:"createdAt"`
}

type ResponseGetProduct struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Category    string `json:"category"`
	ImageUrl    string `json:"imageUrl"`
	Stock       int    `json:"stock"`
	Notes       string `json:"notes"`
	Price       int    `json:"price"`
	Location    string `json:"location"`
	IsAvailable bool   `json:"isAvailable"`
	CreatedAt   string `json:"createdAt"`
}

type ProductUri struct {
	Id string `uri:"id" binding:"required"`
}
