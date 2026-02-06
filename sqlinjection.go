package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	user := os.Args[1]
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", user)

	db, _ := sql.Open("mysql", "user:pass@/dbname")
	db.Query(query)
}
