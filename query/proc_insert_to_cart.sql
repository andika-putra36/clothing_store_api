CREATE OR REPLACE PROCEDURE insert_to_cart(
	p_customer_id	INT
	, p_product_id	INT
)
LANGUAGE plpgsql
AS $$
BEGIN
	INSERT INTO carts(customer_id, product_id)
	VALUES(p_customer_id, p_product_id);
END;
$$;

