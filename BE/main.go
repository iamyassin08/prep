package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/api/routes"
	"github.com/iamyassin08/prep/db"
	_ "github.com/mattn/go-sqlite3"
)

// App holds all application-level dependencies
type App struct {
	db    *db.Queries
	fiber *fiber.App
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize Fiber
	fiberApp := fiber.New(fiber.Config{
		StreamRequestBody: true,
		AppName:           "Prep",
		ServerHeader:      "Fiber",
	})
	// fiberApp.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// }))
	routes.InitPublicRoutes(fiberApp)

	// Initialize routes
	// Initialize database connection
	dbConn, err := sql.Open("sqlite3", "tmp/app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Initialize queries
	queries := db.New(dbConn)
	defer queries.Close()

	// // Create random users
	// if err := app.createRandomUsers(10); err != nil {
	// 	log.Fatal(err)
	// }

	// Start the server
	log.Fatal(fiberApp.Listen(":8080"))
}
func (a *App) createRandomUsers(n int) error {
	firstNames := []string{"John", "Jane", "Alice", "Bob", "Charlie", "Diana", "Eva", "Frank", "Grace", "Henry"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}

	ctx := context.Background()

	for i := 0; i < n; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		email := fmt.Sprintf("%s.%s@example.com", firstName, lastName)

		// Using the CreateUser method from your generated sqlc code
		// Note: You'll need to adjust this based on your actual CreateUser parameters
		_, err := a.db.CreateUser(ctx, db.CreateUserParams{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		})
		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
	}

	return nil
}
