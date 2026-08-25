package transaction

type Service interface {
	GetTransaction(id int) (GetTransactionResponse, error)
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
