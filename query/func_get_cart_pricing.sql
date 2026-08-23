/*
	SELECT * FROM get_cart_pricing(1)
*/

CREATE OR REPLACE FUNCTION get_cart_pricing(
	p_customer_id INT
)
RETURNS TABLE(
	subtotal 			NUMERIC(12, 2)
	, application_fee	NUMERIC(12, 2)
	, total				NUMERIC(12, 2)
)
AS $$
	SELECT
		COALESCE(SUM(products.price), 0)
		, 5000
		, COALESCE((SUM(products.price) + 5000), 0)
	FROM 
		carts
		LEFT JOIN products
			ON products.id = carts.product_id
		LEFT JOIN product_categories 
			ON product_categories.id = products.product_category_id
	WHERE
		products.is_delete = false;
$$ LANGUAGE sql
