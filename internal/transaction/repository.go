package transaction

import "gorm.io/gorm"

type Repository interface {
	GetTransactionHeader(id int) (GetTransactionHeaderResponse, error)
	GetTransactionProducts(id int) ([]GetTransactionProductsResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransactionHeader(id int) (GetTransactionHeaderResponse, error) {
	var response GetTransactionHeaderResponse

	err := r.db.Raw(
		`SELECT * FROM get_transaction_header(?)`,
		id,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) GetTransactionProducts(id int) ([]GetTransactionProductsResponse, error) {
	var response []GetTransactionProductsResponse

	err := r.db.Raw(
		`SELECT * FROM get_transaction_products(?)`,
		id,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}
