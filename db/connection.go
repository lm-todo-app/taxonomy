package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InitDatabase() {
	fmt.Println("Connecting to database")

	connStr := os.Getenv("DB_PATH")
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)
}
