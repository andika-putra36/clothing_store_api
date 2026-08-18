CREATE OR REPLACE PROCEDURE delete_product(
	p_id INT
)
LANGUAGE plpgsql
AS $$
BEGIN
	DELETE FROM products
	WHERE id = p_id;
END;
$$;