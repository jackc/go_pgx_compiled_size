package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var n int
	err = db.QueryRow("select 1").Scan(&n)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Query failed")
		os.Exit(1)
	}

	fmt.Println(n)
}
