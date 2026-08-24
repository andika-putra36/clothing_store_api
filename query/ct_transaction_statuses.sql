CREATE TABLE transaction_statuses(
	id SERIAL PRIMARY KEY
	, transaction_status_category_id INT NOT NULL REFERENCES transaction_status_categories(id)
	, name VARCHAR(255) NOT NULL UNIQUE
	, created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	, updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)