package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
  gde "github.com/joho/godotenv"
)

var DB *sql.DB

func init() {

  log.Println("Loading .env files...")
  if err := gde.Load(".env"); err != nil {
    log.Fatal("Failed to load .env file")
  }

  log.Println("Initializing database... ")
  db, err := connectWithPq()
  if err != nil {
    log.Fatalf("Failed to connect to db: %v", err)
  }
  defer db.Close()
  
  log.Println("Successfully connected to the database")

}

func connectWithPq() (*sql.DB, error) {
  var connString string
  if os.Getenv("ENV") == "development" {
   connString = os.Getenv("LOCAL_DB_URL")
  } else {
    connString = os.Getenv("RENDER_DB_URL")
  }

  db, err := sql.Open("postgres", connString)
  if err != nil {
    return nil, err
  }

  // Configure connection pooling
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(5)
  db.SetConnMaxLifetime(time.Hour)

  // Test the connection
  err = db.Ping()
  if err != nil {
    return nil, err
  }

  return db, nil
}
