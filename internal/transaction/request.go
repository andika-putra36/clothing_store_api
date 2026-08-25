package transaction

type InsertTransactionRequest struct {
	Products       []InsertTransactionProduct `json:"products"`
	ApplicationFee float64                    `json:"application_fee"`
}

type InsertTransactionProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// type GetTransactionRequest struct {
// 	ID int `json:"id"`
// }
