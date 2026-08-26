package transaction

import "errors"

type Service interface {
	GetTransaction(id int) (GetTransactionResponse, error)
	CancelTransaction(id int) error
	AcceptTransaction(id int) error
	RejectTransaction(id int) error
	CompleteTransaction(id int) error
	InsertTransaction(customer_id int, input InsertTransactionRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransaction(id int) (GetTransactionResponse, error) {
	var transactionHeader GetTransactionHeaderResponse

	transactionHeader, err := s.repository.GetTransactionHeader(id)
	if err != nil {
		return GetTransactionResponse{}, err
	}

	var transactionProducts []GetTransactionProductsResponse

	transactionProducts, err = s.repository.GetTransactionProducts(id)
	if err != nil {
		return GetTransactionResponse{}, err
	}

	return GetTransactionResponse{
		Header:   transactionHeader,
		Products: transactionProducts,
	}, nil
}

func (s *service) CancelTransaction(id int) error {
	return s.repository.CancelTransaction(id)
}

func (s *service) AcceptTransaction(id int) error {
	return s.repository.AcceptTransaction(id)
}

func (s *service) RejectTransaction(id int) error {
	return s.repository.RejectTransaction(id)
}

func (s *service) CompleteTransaction(id int) error {
	return s.repository.CompleteTransaction(id)
}

func (s *service) InsertTransaction(customer_id int, input InsertTransactionRequest) error {
	if len(input.ProductIDs) == 0 {
		return errors.New("no products provided")
	}

	if len(input.ProductIDs) != len(input.ProductNames) ||
		len(input.ProductIDs) != len(input.ProductDescriptions) ||
		len(input.ProductIDs) != len(input.ProductPrices) {
		return errors.New("mismatched array lengths")
	}

	return s.repository.InsertTransaction(customer_id, input)
}
