CREATE OR REPLACE PROCEDURE cancel_transaction(
	p_id INT
)
AS $$
BEGIN
	UPDATE transactions
	SET
		transaction_status_id = 4
		, cancelled_at = NOW()
	WHERE id = p_id;
END;
$$ LANGUAGE plpgsql