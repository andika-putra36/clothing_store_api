package cart

type Service interface {
	InsertToCart(customer_id int, input InsertToCartRequest) error
	DeleteFromCart(customer_id int, product_id int) error
	GetCart(customer_id int) (GetCartResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InsertToCart(customer_id int, input InsertToCartRequest) error {
	return s.repository.InsertToCart(customer_id, input)
}

func (s *service) DeleteFromCart(customer_id int, product_id int) error {
	return s.repository.DeleteFromCart(customer_id, product_id)
}

func (s *service) GetCart(customer_id int) (GetCartResponse, error) {
	var cartPricing GetCartPricingResponse

	cartPricing, err := s.repository.GetCartPricing(customer_id)
	if err != nil {
		return GetCartResponse{}, err
	}

	var cartProducts []GetCartProductsResponse

	cartProducts, err = s.repository.GetCartProducts(customer_id)
	if err != nil {
		return GetCartResponse{}, err
	}

	return GetCartResponse{
		Pricing:  cartPricing,
		Products: cartProducts,
	}, nil
}
