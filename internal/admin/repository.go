package admin

import "gorm.io/gorm"

type Repository interface {
	InsertProduct(input InsertProductRequest) error
	UpdateProduct(id int, input UpdateProductRequest) error
	DeleteProduct(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertProduct(input InsertProductRequest) error {
	err := r.db.Exec(
		`CALL insert_product(?, ?, ?, ?, ?)`,
		input.Name,
		input.Description,
		input.Price,
		input.ProductCategoryID,
		input.IsAvailable,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateProduct(id int, input UpdateProductRequest) error {
	err := r.db.Exec(
		`CALL update_product(?, ?, ?, ?, ?, ?)`,
		id,
		input.Name,
		input.Description,
		input.Price,
		input.ProductCategoryID,
		input.IsAvailable,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteProduct(id int) error {
	err := r.db.Exec(
		`CALL delete_product(?)`,
		id,
	).Error
	if err != nil {
		return err
	}
	return nil
}
