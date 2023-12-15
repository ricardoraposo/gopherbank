package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Store struct {
	db *sql.DB
}

var (
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
	dbport   = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() *Store {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, dbport, dbname))
	if err != nil {
		log.Fatal(err)
	}

	return &Store{db}
}

func (s *Store) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintln("Database down: ", err))
	}

	return map[string]string{
		"message": "Db good",
	}
}
