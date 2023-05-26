package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConn() (db *sql.DB) {
	errLoadEnv := godotenv.Load("local.env")
	if errLoadEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errLoadEnv)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	if err != nil {
		// panic(err.Error())
		fmt.Println(err.Error())
	}
	return db
}
