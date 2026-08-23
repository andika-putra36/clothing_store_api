package cart

type InsertToCartRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
	ProductID  int `json:"product_id" binding:"required"`
}

type DeleteFromCartRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
	ProductID  int `json:"product_id" binding:"required"`
}

type GetCartProductsRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
}

type GetCartPricingRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
}
