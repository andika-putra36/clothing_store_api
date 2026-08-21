/*
	SELECT * FROM get_cart_products(1)
*/

CREATE OR REPLACE FUNCTION get_cart_products(
	p_customer_id INT
)
RETURNS TABLE(
	product_id				INT
	, product_category_id 	INT
	, product_category_name VARCHAR(255)
	, name 					VARCHAR(255)
	, description 			TEXT
	, price 				NUMERIC(12, 2)
	, is_available 			BOOLEAN
	, is_delete 			BOOLEAN
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
	FROM 
		carts
		LEFT JOIN products
			ON products.id = carts.product_id
		LEFT JOIN product_categories 
			ON product_categories.id = products.product_category_id
	WHERE
		products.is_delete = false;
$$ LANGUAGE sql
