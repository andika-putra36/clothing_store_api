package cart

import "gorm.io/gorm"

type Repository interface {
	InsertToCart(input InsertToCartRequest) error
	DeleteFromCart(input DeleteFromCartRequest) error
	GetCartProducts(input GetCartProductsRequest) ([]GetCartProductsResponse, error)
	GetCartPricing(input GetCartPricingRequest) (GetCartPricingResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertToCart(input InsertToCartRequest) error {
	err := r.db.Exec(
		`CALL insert_to_cart(?, ?)`,
		input.CustomerID,
		input.ProductID,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteFromCart(input DeleteFromCartRequest) error {
	err := r.db.Exec(
		`CALL delete_from_cart(?, ?)`,
		input.CustomerID,
		input.ProductID,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCartProducts(input GetCartProductsRequest) ([]GetCartProductsResponse, error) {
	var response []GetCartProductsResponse

	err := r.db.Raw(
		`SELECT * FROM get_cart_products(?)`,
		input.CustomerID,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) GetCartPricing(input GetCartPricingRequest) (GetCartPricingResponse, error) {
	var response GetCartPricingResponse

	err := r.db.Raw(
		`SELECT * FROM get_cart_pricing(?)`,
		input.CustomerID,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}
