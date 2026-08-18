CREATE OR REPLACE PROCEDURE insert_product(
	p_name VARCHAR(255)
	, p_description TEXT
	, p_price NUMERIC(12,2)
	, p_product_category_id INT
	, p_is_available BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	INSERT INTO products(product_category_id, name, description, price, is_available)
	VALUES(p_product_category_id, p_name, p_description, p_price, p_is_available);
END;
$$;