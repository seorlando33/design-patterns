package singleton

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDatabase() *sql.DB {

	once.Do(func() {
		conn, err := sql.Open("postgres", getStrConnection())
		if err != nil {
			panic(err)
		}

		err = conn.Ping()
		if err != nil {
			panic(err)
		}

		db = conn
	})

	return db
}

func getStrConnection() string {

	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("user=%s dbname=%s host=%s port=%v password=%s sslmode=disable", user, dbName, host, port, password)
}