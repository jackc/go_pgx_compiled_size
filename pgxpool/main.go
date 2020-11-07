package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	var n int
	err = pool.QueryRow(context.Background(), "select 1").Scan(&n)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Query failed")
		os.Exit(1)
	}

	fmt.Println(n)
}
