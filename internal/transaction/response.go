package transaction

import "time"

type GetTransactionHeaderResponse struct {
	ID                            int       `json:"id"`
	TransactionStatusID           int       `json:"transaction_status_id"`
	TransactionStatusName         string    `json:"transaction_status_name"`
	TransactionStatusCategoryID   int       `json:"transaction_status_category_id"`
	TransactionStatusCategoryName string    `json:"transaction_status_category_name"`
	CustomerID                    int       `json:"customer_id"`
	CustomerName                  string    `json:"customer_name"`
	CreatedAt                     time.Time `json:"created_at"`
	AcceptedAt                    time.Time `json:"accepted_at"`
	CompletedAt                   time.Time `json:"completed_at"`
	CancelledAt                   time.Time `json:"cancelled_at"`
	RejectedAt                    time.Time `json:"rejected_at"`
	SubtotalPrice                 float64   `json:"subtotal_price"`
	ApplicationFee                float64   `json:"application_fee"`
	TotalPrice                    float64   `json:"total_price"`
}

type GetTransactionProductsResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetTransactionResponse struct {
	Header   GetTransactionHeaderResponse     `json:"header"`
	Products []GetTransactionProductsResponse `json:"products"`
}
