package transaction

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Repository interface {
	GetTransactionHeader(id int) (GetTransactionHeaderResponse, error)
	GetTransactionProducts(id int) ([]GetTransactionProductsResponse, error)
	CancelTransaction(id int) error
	AcceptTransaction(id int) error
	RejectTransaction(id int) error
	CompleteTransaction(id int) error
	InsertTransaction(customer_id int, input InsertTransactionRequest) error
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

func (r *repository) CancelTransaction(id int) error {
	err := r.db.Exec(
		`CALL cancel_transaction(?)`,
		id,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) AcceptTransaction(id int) error {
	err := r.db.Exec(
		`CALL accept_transaction(?)`,
		id,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) RejectTransaction(id int) error {
	err := r.db.Exec(
		`CALL reject_transaction(?)`,
		id,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CompleteTransaction(id int) error {
	err := r.db.Exec(
		`CALL complete_transaction(?)`,
		id,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) InsertTransaction(customer_id int, input InsertTransactionRequest) error {
	err := r.db.Exec(
		`CALL insert_transaction(?, ?, ?, ?, ?, ?)`,
		customer_id,
		pq.Array(input.ProductIDs),
		pq.Array(input.ProductNames),
		pq.Array(input.ProductDescriptions),
		pq.Array(input.ProductPrices),
		input.ApplicationFee,
	).Error
	if err != nil {
		return err
	}
	return nil
}
