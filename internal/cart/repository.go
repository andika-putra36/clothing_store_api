package cart

import "gorm.io/gorm"

type Repository interface {
	InsertToCart(customer_id int, input InsertToCartRequest) error
	DeleteFromCart(customer_id int, product_id int) error
	GetCartProducts(customer_id int) ([]GetCartProductsResponse, error)
	GetCartPricing(customer_id int) (GetCartPricingResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertToCart(customer_id int, input InsertToCartRequest) error {
	err := r.db.Exec(
		`CALL insert_to_cart(?, ?)`,
		customer_id,
		input.ProductID,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteFromCart(customer_id int, product_id int) error {
	err := r.db.Exec(
		`CALL delete_from_cart(?, ?)`,
		customer_id,
		product_id,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCartProducts(customer_id int) ([]GetCartProductsResponse, error) {
	var response []GetCartProductsResponse

	err := r.db.Raw(
		`SELECT * FROM get_cart_products(?)`,
		customer_id,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) GetCartPricing(customer_id int) (GetCartPricingResponse, error) {
	var response GetCartPricingResponse

	err := r.db.Raw(
		`SELECT * FROM get_cart_pricing(?)`,
		customer_id,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}
