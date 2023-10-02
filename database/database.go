package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"os"
	"time"
)

// GlobalDB is a global db object that will be used across different packages
// var GlobalDB *gorm.DB
//type Store struct {
//	db *reform.DB
//}

var GlobalDB *reform.DB

// InitDatabase creates a mysql db connection and stores it in the GlobalDB variable
// It reads the environment variables from the .env file and uses them to create the connection
// It returns an error if the connection fails
func InitDatabase() (err error) {
	// Read the environment variables from the .env file
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	// Create the data source name (DSN) using the environment variables
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DATABASE_HOST"],
		config["DATABASE_PORT"],
		config["DB_DATABASE"],
	)
	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	// Create the connection and store it in the GlobalDB variable
	db, err := GetDB(connectionString)
	reformDb := reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))
	GlobalDB = reformDb
	return
}

func GetDB(uri string) (*sql.DB, error) {
	DB, err := PgxCreateDB(uri)
	if err != nil {
		return nil, err
	}
	DB.SetMaxIdleConns(2)
	DB.SetMaxOpenConns(4)
	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)
	return DB, nil
}

func PgxCreateDB(uri string) (*sql.DB, error) {
	connConfig, _ := pgx.ParseConfig(uri)
	afterConnect := stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, `
			 CREATE TABLE IF NOT EXISTS users(
			 	id SERIAL,
				name varchar NOT NULL,
				age int,
				email varchar NOT NULL UNIQUE,
				password varchar NOT NULL
			 );
		`)
		if err != nil {
			return err
		}
		return nil
	})
	pgxdb := stdlib.OpenDB(*connConfig, afterConnect)
	return pgxdb, nil
}
