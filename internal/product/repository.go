package product

import "gorm.io/gorm"

type Repository interface {
	InsertProduct(input InsertProductRequest) error
	UpdateProduct(id int, input UpdateProductRequest) error
	DeleteProduct(id int) error
	GetProducts() ([]GetProductsResponse, error)
	GetProduct(id int) (GetProductResponse, error)
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

func (r *repository) GetProducts() ([]GetProductsResponse, error) {
	var response []GetProductsResponse

	err := r.db.Raw(
		`SELECT * FROM get_products()`,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) GetProduct(id int) (GetProductResponse, error) {
	var response GetProductResponse

	err := r.db.Raw(
		`SELECT * FROM get_product(?)`,
		id,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}
