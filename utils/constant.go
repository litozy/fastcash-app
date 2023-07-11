package utils

//untuk naruh query

//QUERY USER REPOSITORY
const (
	SELECT_ALL_USER = "SELECT id, user_name,  FROM user_credential"
	INSERT_USER = "INSERT INTO user_credential(username, password, created_at, updated_at) VALUES($1, $2, $3, $4)"
	SELECT_USER_BY_NAME = "SELECT id, user_name, password, created_at, updated_at FROM user_credential WHERE username = $1"

	INSERT_TRANSACTION_APPLICATION = "INSERT INTO tx_application(customer_id, loan_product_id, amount, ojk_status_id, date_approval) VALUES($1, $2, $3, $4, $5)"

	GET_ALL_LOAN_PRODUCT   = "SELECT id, product_name, tenor, max_loan, interest, late_interest FROM loan_product ORDER BY id ASC"
	GET_LOAN_PRODUCT_BY_ID = "SELECT id, product_name, tenor, max_loan, interest, late_interest FROM loan_product WHERE id = $1"
	INSERT_LOAN_PRODUCT    = "INSERT INTO loan_product (product_name, tenor, max_loan, interest, late_interest) VALUES ($1, $2, $3, $4, $5)"
	DELETE_LOAN_PRODUCT    = "DELETE FROM loan_product WHERE id=$1)"
	UPDATE_LOAN_PRODUCT    = "UPDATE loan_product SET product_name = $2, tenor = $3, max_loan = $4, interest = $5, late_interest = $6 WHERE id = $1"

	GET_ALL_OJK_STATUS   = "SELECT id, status FROM ojk_status ORDER BY id ASC"
	GET_OJK_STATUS_BY_ID = "SELECT id, status FROM ojk_status WHERE id = $1"
	INSERT_OJK_STATUS    = "INSERT INTO ojk_status (status) VALUES ($1)"
)

