package queries

import "fmt"

func UserTableExists() string {
	query := checkIfTableExists("users")
	return query
}

func CreateUserTable() string {
	query := get("createUserTable")
	return query
}

func GetUserByName(username string) string {
	baseQuery := get("selectAllUserByName")
	query := fmt.Sprintf(baseQuery, username)
	return query
}

func InsertUser(username, passhash string) string {
	baseQuery := get("insertUser")
	query := fmt.Sprintf(baseQuery, username, passhash)
	return query
}
