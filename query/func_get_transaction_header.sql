CREATE OR REPLACE FUNCTION get_transaction_header(
	p_id INT
)
RETURNS TABLE(
	id 									INT
	, transaction_status_id 			INT
	, transaction_status_name 			VARCHAR(255)
	, transaction_status_category_id	INT
	, transaction_status_category_name	VARCHAR(255)
	, customer_id						INT
	, customer_name						VARCHAR(255)
	, created_at						TIMESTAMPTZ
	, accepted_at						TIMESTAMPTZ
	, completed_at						TIMESTAMPTZ
	, cancelled_at						TIMESTAMPTZ
	, rejected_at						TIMESTAMPTZ
	, subtotal_price					NUMERIC(12, 2)
	, application_fee					NUMERIC(12, 2)
	, total_price						NUMERIC(12, 2)
)
AS $$
	SELECT 
		transactions.id
		, transactions.transaction_status_id
		, transaction_statuses.name
		, transaction_statuses.transaction_status_category_id
		, transaction_status_categories.name
		, transactions.customer_id
		, customers.first_name
		, transactions.created_at
		, transactions.accepted_at
		, transactions.completed_at
		, transactions.cancelled_at
		, transactions.rejected_at
		, transactions.subtotal_price
		, transactions.application_fee
		, transactions.total_price
	FROM 
		transactions
		LEFT JOIN transaction_statuses
			ON transaction_statuses.id = transactions.transaction_status_id
		LEFT JOIN transaction_status_categories
			ON transaction_status_categories.id = transaction_statuses.transaction_status_category_id
		LEFT JOIN customers
			ON customers.id = transactions.customer_id
	WHERE 
		transactions.id = p_id;
$$ LANGUAGE sql