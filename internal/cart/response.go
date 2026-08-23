package cart

type GetCartProductsResponse struct {
	ID                  int     `json:"id"`
	ProductCategoryID   int     `json:"product_category_id"`
	ProductCategoryName string  `json:"product_category_name"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Price               float64 `json:"price"`
}

type GetCartPricingResponse struct {
	Subtotal       float64 `json:"subtotal"`
	ApplicationFee float64 `json:"application_fee"`
	Total          float64 `json:"total"`
}
