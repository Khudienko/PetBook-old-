package driver

import (
	//"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func init() {
	gotenv.Load()
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ConnectDB() *sqlx.DB {
	var err error
	var db *sqlx.DB
	host := os.Getenv("HOST_POSTGRES")
	port := os.Getenv("PORT_POSTGRES")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sqlx.Open("postgres", connStr)
	logErr(err)
	err = db.Ping()
	if err != nil {
		logErr(fmt.Errorf("can't ping, err: %s", err.Error()))
	}
	return db
}
