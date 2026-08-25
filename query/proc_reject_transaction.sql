CREATE OR REPLACE PROCEDURE reject_transaction(
	p_id INT
)
AS $$
BEGIN
	UPDATE transactions
	SET
		transaction_status_id = 5
		, rejected_at = NOW()
	WHERE id = p_id;
END;
$$ LANGUAGE plpgsql