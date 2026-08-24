CREATE TABLE transactions(
	id SERIAL PRIMARY KEY
	, transaction_status_id INT NOT NULL REFERENCES transaction_statuses(id)
	, created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	, accepted_at TIMESTAMPTZ
	, completed_at TIMESTAMPTZ
	, cancelled_at TIMESTAMPTZ
	, rejected_at TIMESTAMPTZ
	, subtotal_price NUMERIC(12, 2) NOT NULL DEFAULT 0
	, application_fee NUMERIC(12, 2) NOT NULL DEFAULT 0
	, total_price NUMERIC(12, 2) NOT NULL DEFAULT 0
)