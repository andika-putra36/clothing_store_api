/*
	SELECT * FROM get_login_credential('andika@mail.test');
*/

CREATE OR REPLACE FUNCTION get_login_credential(
	p_email VARCHAR(255)
)
RETURNS TABLE (
	user_id 		INT
	, role_id 		INT
	, email 		VARCHAR(255)
	, password_hash TEXT
	, customer_id	INT
	, admin_id		INT
)
AS $$
	SELECT 
		users.id
		, users.role_id
		, users.email
		, users.password_hash
		, customers.id
		, admins.id
	FROM 
		users
		LEFT JOIN customers
			ON customers.user_id = users.id
		LEFT JOIN admins
			ON admins.user_id = users.id
	WHERE 
		email = p_email;
$$ LANGUAGE sql