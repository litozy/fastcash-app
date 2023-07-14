package utils

//untuk naruh query

//QUERY USER REPOSITORY
const (
	SELECT_ALL_USER     = "SELECT id, user_name,  FROM user_credential"
	INSERT_USER         = "INSERT INTO user_credential(username, password, created_at, updated_at) VALUES($1, $2, $3, $4)"
	SELECT_USER_BY_NAME = "SELECT id, user_name, password, created_at, updated_at FROM user_credential WHERE username = $1"
	SELECT_USER_BY_ID = "SELECT id, user_name, password, created_at, updated_at FROM user_credential WHERE id = $1"

	INSERT_TRANSACTION_APPLICATION = `
	INSERT INTO tx_application(customer_id, loan_product_id, amount, ojk_status_id, date_approval, create_by, update_by) 
	SELECT $1, $2, $3, $4, $5, name, $6
	FROM customer
	WHERE id = $7
	`
	GET_ALL_TRANSACTION_APPLICATION = `
	SELECT tx.id AS id, tx.customer_id AS custId, c.name AS custName, c.bank_account AS bankaccount, c.nik AS nik, p.tenor AS tenor, tx.amount AS amount, tx.date_approval, o.status AS status, o.description AS description 
	FROM tx_application AS tx
	JOIN customer AS c ON tx.customer_id = c.id
	JOIN ojk_status AS o ON tx.ojk_status_id = o.id 
	JOIN loan_product AS p ON tx.loan_product_id = p.id
	ORDER BY id ASC
	`
	GET_TRANSACTION_APPLICATION_BY_ID = `
	SELECT tx.id AS id, tx.customer_id AS custId, c.name AS custName, c.bank_account AS bankaccount, c.nik AS nik, p.tenor AS tenor, tx.amount AS amount, tx.date_approval, o.status AS status, o.description AS description 
	FROM tx_application AS tx
	JOIN customer AS c ON tx.customer_id = c.id
	JOIN ojk_status AS o ON tx.ojk_status_id = o.id 
	JOIN loan_product AS p ON tx.loan_product_id = p.id
	WHERE tx.id = $1
	`
	UPDATE_OJK_STATUS_TRANSACTION_APPLICATION = "UPDATE tx_application SET ojk_status_id = $1, update_by = $2 WHERE id = $3"
	UPDATE_OJK_STATUS_DATE_APPROVAL           = "UPDATE tx_application SET date_approval = $1 WHERE id = $2"

	INSERT_TRANSACTION_PAYMENT = `
	INSERT INTO tx_payment(aplication_id, payment, created_by, created_at, status)
	SELECT $1, $2, create_by, $3, $4
	FROM tx_application
	WHERE id = $5
	`
	GET_TRANSACTION_PAYMENT_BY_ID = `
	SELECT tx.customer_id AS custId, c.name AS custName, CONCAT(p.tenor, ' bulan') AS tenor, (tx.amount) * (p.interest / 100) + (tx.amount) AS MustToPay, 
    COALESCE(SUM(tp.payment), 0) AS paid, 
    tx.amount - COALESCE(SUM(tp.payment), 0) AS needtopay, (tx.amount / p.tenor) * (p.interest / 100) + (tx.amount / p.tenor) AS NeedToPay1Month ,
    (DATE_TRUNC('month', tx.date_approval) + INTERVAL '25 days' + INTERVAL '1 month' * (p.tenor)) AS deadlinepayment
FROM tx_application AS tx
LEFT JOIN tx_payment AS tp ON tx.id = tp.aplication_id
JOIN customer AS c ON tx.customer_id = c.id
JOIN loan_product AS p ON tx.loan_product_id = p.id
WHERE tx.id = $1
GROUP BY tx.customer_id, c.name, p.tenor, tx.amount, p.interest, tx.date_approval;
	`
	GET_TRANSACTION_PAYMENT_BY_ID_VALIDATE = `
	SELECT tx.customer_id AS custId, c.name AS custName, CONCAT(p.tenor, ' bulan') AS tenor, (tx.amount) * (p.interest / 100) + (tx.amount) AS MustToPay, 
    COALESCE(SUM(tp.payment), 0) AS paid, 
    tx.amount - COALESCE(SUM(tp.payment), 0) AS needtopay, (tx.amount / p.tenor) * (p.interest / 100) + (tx.amount / p.tenor) AS OneMonthPayment
FROM tx_application AS tx
LEFT JOIN tx_payment AS tp ON tx.id = tp.aplication_id
JOIN customer AS c ON tx.customer_id = c.id
JOIN loan_product AS p ON tx.loan_product_id = p.id
WHERE tx.id = $1
GROUP BY tx.customer_id, c.name, p.tenor, tx.amount, p.interest;
	`
	UPDATE_TRANSACTION_PAYMENT_STATUS = "UPDATE tx_payment SET status = $1 WHERE id = $2"

	GET_ALL_CUSTOMER   = "SELECT id, user_id, name, address, nik, birthdate, family_member, family_phone, family_address, status FROM customer ORDER BY id ASC"
	GET_CUSTOMER_BY_ID = "SELECT id, user_id, name, address, nik, birthdate, family_member, family_phone, family_address, status FROM customer WHERE id = $1"
	INSERT_CUSTOMER    = "INSERT INTO customer (user_id, name, address, nik, birthdate, family_member, family_phone, family_address, status, bank_account) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	DELETE_CUSTOMER    = "DELETE FROM customer WHERE id=$1)"
	UPDATE_CUSTOMER    = "UPDATE customer SET user_id = $2, name = $3, address = $4, nik = $5, birthdate = $6, family_member = $7, family_phone = $8, family_address = $9, bank_account = $10 WHERE id = $1"
	UPDATE_CUSTOMER_STATUS = "UPDATE customer SET status = $1 WHERE id = $2"

	GET_ALL_LOAN_PRODUCT   = "SELECT id, product_name, tenor, max_loan, interest, late_interest FROM loan_product ORDER BY id ASC"
	GET_LOAN_PRODUCT_BY_ID = "SELECT id, product_name, tenor, max_loan, interest, late_interest FROM loan_product WHERE id = $1"
	INSERT_LOAN_PRODUCT    = "INSERT INTO loan_product (product_name, tenor, max_loan, interest, late_interest) VALUES ($1, $2, $3, $4, $5)"
	DELETE_LOAN_PRODUCT    = "DELETE FROM loan_product WHERE id=$1)"
	UPDATE_LOAN_PRODUCT    = "UPDATE loan_product SET product_name = $2, tenor = $3, max_loan = $4, interest = $5, late_interest = $6 WHERE id = $1"

	GET_ALL_OJK_STATUS   = "SELECT id, status, description FROM ojk_status ORDER BY id ASC"
	GET_OJK_STATUS_BY_ID = "SELECT id, status, description FROM ojk_status WHERE id = $1"
	INSERT_OJK_STATUS    = "INSERT INTO ojk_status (status, description) VALUES ($1, $2)"
	DELETE_OJK_STATUS    = "DELETE FROM ojk_status WHERE id =$1 "
	UPDATE_OJK_STATUS    = "UPDATE ojk_status SET status = $1, description $2 WHERE id = $3 "

)
