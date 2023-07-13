SELECT * FROM customer;
SELECT * FROM loan_product;v
SELECT * FROM ojk_status;v
SELECT * FROM tx_application;
SELECT * FROM tx_payment;
SELECT * FROM user_credential;

CREATE TABLE user_credential (
id SERIAL PRIMARY KEY NOT NULL,
	username VARCHAR(50),
	password VARCHAR (50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updates_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE customer (
id SERIAL PRIMARY KEY NOT NULL,
user_id INT NOT NULL,
name VARCHAR (50) NOT NULL,
address VARCHAR(50) NOT NULL,
nik INT NOT NULL,
birthdate DATE NOT NULL,
family_member VARCHAR(50) NOT NULL,
family_phone VARCHAR NOT NULL,
family_address VARCHAR (50) NOT NULL,
status VARCHAR (20) NOT NULL,
FOREIGN KEY (user_id) REFERENCES user_credential(id));

CREATE TABLE loan_product (
id SERIAL PRIMARY KEY NOT NULL,
product_name VARCHAR(50),
tenor INT NOT NULL,
max_loan FLOAT,
interest FLOAT,
late_interest FLOAT);

CREATE TABLE ojk_status (
id SERIAL PRIMARY KEY NOT NULL,
status INT NOT NULL
description VARCHAR (50) NOT NULL);

CREATE TABLE tx_application (
id SERIAL PRIMARY KEY NOT NULL,
customer_id INT NOT NULL,
loan_product_id INT NOT NULL,
amount FLOAT,
ojk_status_id INT NOT NULL,
date_approval DATE,
create_by VARCHAR (50),
update_by VARCHAR (50),
FOREIGN KEY (customer_id) REFERENCES customer(id),
FOREIGN KEY (loan_product_id) REFERENCES loan_product(id),
FOREIGN KEY (ojk_status_id) REFERENCES ojk_status(id));

CREATE TABLE tx_payment (
id SERIAL PRIMARY KEY NOT NULL,
payment FLOAT,
aplication_id INT NOT NULL,
created_by VARCHAR,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (aplication_id) REFERENCES tx_application(id));

