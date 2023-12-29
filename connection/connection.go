package ksi

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connection() {
	// Replace the connection parameters with your actual database details
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	// Perform database operations (query, insert, update, etc.) as needed
	// ...

	// Example query
	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the query results
	for rows.Next() {
		var col1, col2 string
		err := rows.Scan(&col1, &col2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(col1, col2)
	}

	// Handle any errors that occurred during the query
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
