CREATE TABLE transaction_products(
	transaction_id INT NOT NULL REFERENCES transactions(id)
	, product_id INT NOT NULL REFERENCES products(id)
	, product_name VARCHAR(255)
	, product_description TEXT
	, product_price NUMERIC(12, 2)
	, PRIMARY KEY(transaction_id, product_id)
)