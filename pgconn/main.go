package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgconn"
)

func main() {
	pgConn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "pgconn failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer pgConn.Close(context.Background())

	result := pgConn.ExecParams(context.Background(), "SELECT 1", nil, nil, nil, nil)
	for result.NextRow() {
		fmt.Println(string(result.Values()[0]))
	}
	_, err = result.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed reading result: %v\n", err)
		os.Exit(1)
	}
}
