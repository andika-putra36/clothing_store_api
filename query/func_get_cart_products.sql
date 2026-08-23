/*
	SELECT * FROM get_cart_products(1)
*/

CREATE OR REPLACE FUNCTION get_cart_products(
	p_customer_id INT
)
RETURNS TABLE(
	id						INT
	, product_category_id 	INT
	, product_category_name VARCHAR(255)
	, name 					VARCHAR(255)
	, description 			TEXT
	, price 				NUMERIC(12, 2)
)
AS $$
	SELECT
		products.id
		, products.product_category_id
		, product_categories.name
		, products.name
		, products.description
		, products.price
	FROM 
		carts
		LEFT JOIN products
			ON products.id = carts.product_id
		LEFT JOIN product_categories 
			ON product_categories.id = products.product_category_id
	WHERE
		products.is_available = true
		AND products.is_delete = false;
$$ LANGUAGE sql
