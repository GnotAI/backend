package db

import (
  "os"
  "log"
  "context"

  "github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Initdb() error {

  var dbURL string
  ctx := context.Background()

  ENV := os.Getenv("ENV")
  if ENV == "production" {
    dbURL = os.Getenv("RENDER_DB_URL")
  } else {
    dbURL = os.Getenv("LOCAL_DB_URL")
  }

  // Connect to db
  Pool, err := pgxpool.New(ctx, dbURL)
  if err != nil {
   log.Fatalf("Unable to connect to database: %v", err)
  }

  // Test connection
  err = Pool.Ping(ctx)
  if err != nil {
    log.Fatalf("Unable to ping to database: %v", err)
  }

  return nil
}
