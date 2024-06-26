// http://localhost:8080/swagger/index.html

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	h "effectiveMobileGo/internal/handle"

	_ "effectiveMobileGo/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

// @title Cars API
// @description This is a sample API for managing cars and their owners.
// @version 1.0
// @host localhost:8080
func main() {
	db := initDB()
	defer db.Close()

	router := mux.NewRouter()

	// GET /cars - Get all cars
	router.HandleFunc("/cars",
		func(w http.ResponseWriter, r *http.Request) { h.GetCars(w, r, db) }).Methods("GET")

	// GET /cars/{filterField}/{valueField} - Get cars by filter
	router.HandleFunc("/cars/{filterField}/{valueField}",
		func(w http.ResponseWriter, r *http.Request) { h.GetCars(w, r, db) }).Methods("GET")

	// DELETE /cars/{id} - Delete a car by ID
	router.HandleFunc("/cars/{id}",
		func(w http.ResponseWriter, r *http.Request) { h.DeleteCars(w, r, db) }).Methods("DELETE")

	// PATCH /cars/{id} - Update a car by ID
	router.HandleFunc("/cars/{id}",
		func(w http.ResponseWriter, r *http.Request) { h.PatchCars(w, r, db) }).Methods("PATCH")

	// POST /cars - Create a new cars
	router.HandleFunc("/cars",
		func(w http.ResponseWriter, r *http.Request) { h.PostCars(w, r, db) }).Methods("POST")

	// Swagger Docs
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("Enter the Postgres port number:")
		fmt.Scan(&port)
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func initDB() *sql.DB {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Print("Enter the Postgres port number:")
		fmt.Scan(&dbPort)
	}

	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), dbPort, os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/../migrations", os.Getenv("PWD")),
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT COUNT( * ) FROM cars")
	var count int
	if err := row.Scan(&count); err != nil {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}

	return db
}
