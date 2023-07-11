package utils

//untuk naruh query

//QUERY USER REPOSITORY
const (
	SELECT_ALL_USER = "SELECT id, user_name,  FROM user_credential"
	INSERT_USER = "INSERT INTO user_credential(username, password, created_at, updated_at) VALUES($1, $2, $3, $4)"
	SELECT_USER_BY_NAME = "SELECT id, user_name, password, created_at, updated_at FROM user_credential WHERE username = $1"

	INSERT_TRANSACTION_APPLICATION = "INSERT INTO tx_application(customer_id, loan_product_id, amount, ojk_status_id, date_approval) VALUES($1, $2, $3, $4, $5)"
)