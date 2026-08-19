package admin

type Service interface {
	InsertProduct(input InsertProductRequest) error
	UpdateProduct(id int, input UpdateProductRequest) error
	DeleteProduct(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InsertProduct(input InsertProductRequest) error {
	return s.repository.InsertProduct(input)
}

func (s *service) UpdateProduct(id int, input UpdateProductRequest) error {
	return s.repository.UpdateProduct(id, input)
}

func (s *service) DeleteProduct(id int) error {
	return s.repository.DeleteProduct(id)
}
