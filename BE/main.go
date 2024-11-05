package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iamyassin08/prep/db" // Adjust this import based on your project structure
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Seed random number generator for generating random users
	rand.Seed(time.Now().UnixNano())

	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
		AppName:           "Prep",
		ServerHeader:      "Fiber",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Connect to SQLite database
	dB, err := sql.Open("sqlite3", "tmp/app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dB.Close()

	// Initialize the DB variable with the connection
	db.DB = db.New(dB)

	// Create random users
	createRandomUsers(dB, 10)

	log.Fatal(app.Listen(":8080"))
}

func createRandomUsers(db *sql.DB, n int) {
	firstNames := []string{"John", "Jane", "Alice", "Bob", "Charlie", "Diana", "Eva", "Frank", "Grace", "Henry"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}

	for i := 0; i < n; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		email := fmt.Sprintf("%s.%s@example.com", firstName, lastName)

		_, err := db.Exec("INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)", firstName, lastName, email)
		if err != nil {
			log.Fatal(err)
		}
	}
}
