package admin

type InsertProductRequest struct {
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description" binding:"required"`
	Price             float64 `json:"price" binding:"required"`
	ProductCategoryID int     `json:"product_category_id" binding:"required"`
	IsAvailable       *bool   `json:"is_available" binding:"required"`
}

type UpdateProductRequest struct {
	ID                int     `json:"id" binding:"required"`
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description" binding:"required"`
	Price             float64 `json:"price" binding:"required"`
	ProductCategoryID int     `json:"product_category_id" binding:"required"`
	IsAvailable       *bool   `json:"is_available" binding:"required"`
}

type DeleteProductRequest struct {
	ID int `json:"id" binding:"required"`
}
