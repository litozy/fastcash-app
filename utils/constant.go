package utils

//untuk naruh query

const (
	GET_ALL_INTEREST   = "SELECT id, interest_rate FROM interest ORDER BY id ASC"
	GET_INTEREST_BY_ID = "SELECT id, interest_rate FROM interest WHERE id = $1"
	INSERT_INTEREST    = "INSERT INTO interest (interest_rate) VALUES ($1)"
	DELETE_INTEREST    = "DELETE FROM interest WHERE id=$1)"
	UPDATE_INTEREST    = "UPDATE interest SET interest_rate = $2 WHERE id = $1"

	GET_ALL_LATE_INTEREST   = "SELECT id, interest_rate FROM late_interest ORDER BY id ASC"
	GET_LATE_INTEREST_BY_ID = "SELECT id, interest_rate FROM late_interest WHERE id = $1"
	INSERT_LATE_INTEREST    = "INSERT INTO late_interest (interest_rate) VALUES ($1)"
	DELETE__LATE_INTEREST   = "DELETE FROM late_interest WHERE id=$1)"
	UPDATE_LATE_INTEREST    = "UPDATE late_interest SET interest_rate = $2 WHERE id = $1"
)
