CREATE OR REPLACE FUNCTION get_refresh_token(
	p_token TEXT
)
RETURNS TABLE(
	id INT
	, user_id INT
	, email VARCHAR(255)
	, role_id INT
	, token TEXT
	, expired_at TIMESTAMP
	, customer_id	INT
	, admin_id		INT
)
AS $$
	SELECT 
		refresh_tokens.id
		, refresh_tokens.user_id
		, users.email
		, users.role_id
		, refresh_tokens.token
		, refresh_tokens.expired_at
		, customers.id
		, admins.id
	FROM
		refresh_tokens
		LEFT JOIN users
			ON users.id = refresh_tokens.user_id
		LEFT JOIN customers
			ON customers.user_id = users.id
		LEFT JOIN admins
			ON admins.user_id = users.id
	WHERE 
		refresh_tokens.token = p_token;
$$ LANGUAGE sql;