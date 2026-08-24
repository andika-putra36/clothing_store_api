CREATE OR REPLACE PROCEDURE insert_transaction(
	p_product_ids				INT[]
	, p_product_names			VARCHAR(255)[]
	, p_product_descriptions	TEXT[]
	, p_product_prices			NUMERIC(12, 2)[]
	, p_application_fee			NUMERIC(12, 2)
)
AS $$
DECLARE
	v_transaction_id INT;
BEGIN
	INSERT INTO transactions(transaction_status_id)
	VALUES(1)
	RETURNING id INTO v_transaction_id;

	INSERT INTO transaction_products(
		product_id
		, product_name
		, product_description
		, product_price
	)
	VALUES(
		UNNEST(p_product_ids)
		, UNNEST(p_product_names)
		, UNNEST(p_product_descriptions)
		, UNNEST(p_product_prices)
	);

	UPDATE transactions t_transaction
	SET
		subtotal_price = t_transaction_products.subtotal_price
		, application_fee = p_application_fee
		, total_price = t_transaction_products.subtotal_price + p_application_fee
	FROM 
		(
			SELECT COALESCE(SUM(product_price), 0) AS subtotal_price
			FROM transaction_products
			WHERE transaction_id = v_transaction_id
		) t_transaction_products
	WHERE t_transaction.id = v_transaction_id;
END;
$$ LANGUAGE plpgsql