CREATE OR REPLACE PROCEDURE complete_transaction(
	p_id INT
)
AS $$
BEGIN
	UPDATE transactions
	SET
		transaction_status_id = 3
		, completed_at = NOW()
	WHERE id = p_id;
END;
$$ LANGUAGE plpgsql