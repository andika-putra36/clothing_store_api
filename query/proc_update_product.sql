CREATE OR REPLACE PROCEDURE update_product(
	p_id 					INT
	, p_name 				VARCHAR(255)
	, p_description 		TEXT
	, p_price 				NUMERIC(12,2)
	, p_product_category_id	INT
	, p_is_available 		BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	UPDATE products
	SET
		product_category_id	= p_product_category_id
		, name 				= p_name
		, description 		= p_description
		, price 			= p_price
		, is_available 		= p_is_available
	WHERE id = p_id;
END;
$$;