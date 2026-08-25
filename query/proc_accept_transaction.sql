CREATE OR REPLACE PROCEDURE accept_transaction(
	p_id INT
)
AS $$
BEGIN
	UPDATE transactions
	SET
		transaction_status_id = 2
		, accepted_at = NOW()
	WHERE id = p_id;
END;
$$ LANGUAGE plpgsql