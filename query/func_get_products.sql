/*
	SELECT * FROM get_products()
*/

CREATE OR REPLACE FUNCTION get_products()
RETURNS TABLE(
	id 						INT
	, product_category_id 	INT
	, product_category_name VARCHAR(255)
	, name 					VARCHAR(255)
	, description 			TEXT
	, price 				NUMERIC(12, 2)
	, is_available 			BOOLEAN
	, is_delete 			BOOLEAN
	, created_at 			TIMESTAMPTZ
	, updated_at 			TIMESTAMPTZ
)
AS $$
	SELECT
		products.id
		, products.product_category_id
		, product_categories.name
		, products.name
		, products.description
		, products.price
		, products.is_available
		, products.is_delete
		, products.created_at
		, products.updated_at
	FROM 
		products
		LEFT JOIN product_categories 
			ON product_categories.id = products.product_category_id
	WHERE
		products.is_delete = false;
$$ LANGUAGE sql