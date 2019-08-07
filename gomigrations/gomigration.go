package gomigrations

import (
	//"database/sql"
	//"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found:%v", err)
	}
}

func Migrate(db *sqlx.DB) (err error) {
	migrations := getMigrations()
	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Printf("can't work gomigrations:%v", err)
	}
	fmt.Printf("Applied %d gomigrations!\n", n)
	err = db.Ping()
	if err != nil {
		log.Printf("can't ping:%v", err)
	}
	return
}
func getMigrations() *migrate.FileMigrationSource {
	return &migrate.FileMigrationSource{
		Dir: "gomigrations/migrations",
	}
}
