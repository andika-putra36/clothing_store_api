package admin

import "gorm.io/gorm"

type Repository interface {
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

func (r *repository) UpdateProduct(input UpdateProductRequest) error {
	err := r.db.Exec(
		`CALL update_product(?, ?, ?, ?, ?, ?)`,
		input.ID,
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

func (r *repository) DeleteProduct(input DeleteProductRequest) error {
	err := r.db.Exec(
		`CALL delete_product(?)`,
		input.ID,
	).Error
	if err != nil {
		return err
	}
	return nil
}
