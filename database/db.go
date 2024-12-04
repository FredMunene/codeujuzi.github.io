package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	var errConn error

	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading environment vars", err)
		return
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	DB, errConn = sql.Open("postgres", connStr)
	if errConn != nil {
		fmt.Println("error opening sql", errConn)
		return
	}

	if err = DB.Ping(); err != nil {
		fmt.Println("error connecting to db", err)
		return
	}

	fmt.Println("Db connected successfully")
}
