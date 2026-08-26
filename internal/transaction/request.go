package transaction

type InsertTransactionRequest struct {
	ProductIDs          []int     `json:"product_ids"`
	ProductNames        []string  `json:"product_names"`
	ProductDescriptions []string  `json:"product_descriptions"`
	ProductPrices       []float64 `json:"product_prices"`
	ApplicationFee      float64   `json:"application_fee"`
}

// type InsertTransactionProduct struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// }

// type GetTransactionRequest struct {
// 	ID int `json:"id"`
// }
