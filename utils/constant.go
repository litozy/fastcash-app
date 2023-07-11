package utils

//untuk naruh query

const (
	GET_ALL_LOAN_PRODUCT   = "SELECT id, product_name, tenor, max_loan, interest, late_interest FROM loan_product ORDER BY id ASC"
	GET_LOAN_PRODUCT_BY_ID = "SELECT id, product_name, tenor, max_loan, interest, late_interest FROM loan_product WHERE id = $1"
	INSERT_LOAN_PRODUCT    = "INSERT INTO loan_product (product_name, tenor, max_loan, interest, late_interest) VALUES ($1, $2, $3, $4, $5)"
	DELETE_LOAN_PRODUCT    = "DELETE FROM loan_product WHERE id=$1)"
	UPDATE_LOAN_PRODUCT    = "UPDATE loan_product SET product_name = $2, tenor = $3, max_loan = $4, interest = $5, late_interest = $6 WHERE id = $1"

	GET_ALL_OJK_STATUS   = "SELECT id, status, description FROM ojk_status ORDER BY id ASC"
	GET_OJK_STATUS_BY_ID = "SELECT id, status, description FROM ojk_status WHERE id = $1"
	INSERT_OJK_STATUS    = "INSERT INTO ojk_status (status, description) VALUES ($1)"
	DELETE_OJK_STATUS    = "DELETE FROM ojk_status WHERE id =$1 "
	UPDATE_OJK_STATUS    = "UPDATE ojk_status SET status = $2, description $3 WHERE id = $1 "

	GET_ALL_TX_LOAN_APPLICATION   = "SELECT id, customer_id, loan_product_id, amount, ojk_status_id, date_approval,created_by, updated_by FROM tx_application ORDER BY ASC"
	GET_TX_LOAN_APPLICATION_BY_ID = "SELECT id, customer_id, loan_product_id, amount, ojk_status_id, date_approval,created_by, updated_by FROM tx_application WHERE id = $1"
	INSERT_TX_LOAN_APPLICATION    = "INSERT INTO tx_application (customer_id, loan_product_id, amount, ojk_status_id, date_approval,created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7)"
)
