package queries

import (
	"fmt"
	"io/ioutil"
)

func checkIfTableExists(tableName string) string {
	query := get("tableExists")
	return fmt.Sprintf(query, tableName)
}

func get(queryName string) string {
	fileContent, _ := ioutil.ReadFile(fmt.Sprintf("./db/queries/%s.sql", queryName))
	return string(fileContent)
}
