package main

import (
	"database/sql"
	"log"
	"movie-api/cmd"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DBURL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}()

	// We configure our own database setting to limit the number of connections.
	// So we won't drain the database connection.
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetConnMaxIdleTime(time.Second * 20)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(200)

	app := echo.New()
	dependencies := cmd.New(db)

	dependencies.MovieHandler(app.Group("/movie"))

	port := os.Getenv("PORT")
	app.Logger.Fatal(app.Start(":" + port))
}
