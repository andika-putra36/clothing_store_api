/*
	CALL register_customer(
		'putra@mail.com'
		, 'passwordHash'
		, 'Putra'
		, 'Testing'
		, '0812--------'
	);
*/

CREATE OR REPLACE PROCEDURE register_customer(
	p_email VARCHAR(255)
	, p_password_hash TEXT
	, p_first_name VARCHAR(255)
	, p_last_name VARCHAR(255)
	, p_phone_number VARCHAR(255)
)
AS $$
DECLARE
	v_user_id INT;
BEGIN	
	INSERT INTO users(role_id, email, password_hash)
	VALUES(
		3
		, p_email
		, p_password_hash
	)
	RETURNING id INTO v_user_id;

	INSERT INTO customers(user_id, first_name, last_name, phone_number)
	VALUES(
		v_user_id
		, p_first_name
		, p_last_name
		, p_phone_number
	);
END;
$$ LANGUAGE plpgsql