CREATE OR REPLACE FUNCTION get_transaction_products(
	p_transaction_id INT
)
RETURNS TABLE(
	product_id 				INT
	, product_name 			VARCHAR(255)
	, product_description	TEXT
	, product_price 		NUMERIC(12, 2)
)
AS $$
	SELECT
		product_id
		, product_name
		, product_description
		, product_price
	FROM 
		transaction_products
	WHERE transaction_products.transaction_id = p_transaction_id;
$$ LANGUAGE sql