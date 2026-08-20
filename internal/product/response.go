package product

import "time"

type GetProductsResponse struct {
	ID                  int       `json:"id"`
	ProductCategoryID   int       `json:"product_category_id"`
	ProductCategoryName string    `json:"product_category_name"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Price               float64   `json:"price"`
	IsAvailable         bool      `json:"is_available"`
	IsDelete            bool      `json:"is_delete"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type GetProductResponse struct {
	ID                  int       `json:"id"`
	ProductCategoryID   int       `json:"product_category_id"`
	ProductCategoryName string    `json:"product_category_name"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Price               float64   `json:"price"`
	IsAvailable         bool      `json:"is_available"`
	IsDelete            bool      `json:"is_delete"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
