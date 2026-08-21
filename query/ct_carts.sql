CREATE TABLE carts(
	customer_id INT NOT NULL REFERENCES customers(id)
	, product_id INT NOT NULL REFERENCES products(id)
	, PRIMARY KEY(customer_id, product_id)
);