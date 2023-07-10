package utils

//untuk naruh query

//QUERY USER REPOSITORY
const (
	SELECT_ALL_USER = "SELECT id, user_name, is_active FROM user_credential"
	INSERT_USER = "INSERT INTO user_credential(id, user_name, password, is_active) VALUES($1, $2, $3, $4)"
	SELECT_USER_BY_NAME = "SELECT id, user_name, password, is_active FROM user_credential WHERE user_name = $1"
	
)