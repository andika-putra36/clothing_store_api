CREATE TABLE products(
	id SERIAL PRIMARY KEY
	, product_category_id INT NOT NULL UNIQUE REFERENCES product_categories(id)
	, name VARCHAR(255) NOT NULL
	, description TEXT
	, price NUMERIC(12,2) NOT NULL DEFAULT 0
	, is_available BOOLEAN DEFAULT TRUE
	, created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	, updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);