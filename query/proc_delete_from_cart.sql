CREATE OR REPLACE PROCEDURE delete_from_cart(
	p_customer_id	INT
	, p_product_id	INT
)
LANGUAGE plpgsql
AS $$
BEGIN
	DELETE FROM carts
	WHERE 
		customer_id = p_customer_id
		AND product_id = p_product_id;
END;
$$;