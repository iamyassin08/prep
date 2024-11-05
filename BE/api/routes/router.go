package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/api/handler"
	"github.com/iamyassin08/prep/api/middlewares"
	"github.com/iamyassin08/prep/docs"
	"github.com/iamyassin08/prep/identity"
	"github.com/iamyassin08/prep/shared"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func InitPublicRoutes(app *fiber.App) {
	// apiHandler := &handler.ApiHandler{queries: db.DB}
	apihandler := handler.ApiHandler{}
	identityManager := identity.NewIdentityManager()
	registerUseCase := shared.NewRegistraterUseCase(identityManager)
	docs.SwaggerInfo.BasePath = "/api/v1"
	public := app.Group("/")
	{
		// Public routes: registration, login, and health check
		public.Post("/api/v1/register", handler.RegisterHandler(registerUseCase)) // Register
		public.Post("/api/v1/login", handler.LoginHandler(registerUseCase))       // Login
		public.Get("/swagger/*", fiberSwagger.WrapHandler)                        // Swagger UI
		public.Get("/api/v1/healthz", handler.HealthCheck)                        // Health check

		// User-related routes (as per the request)
		public.Get("/api/v1/users", apihandler.ListUsers)         // List all users
		public.Get("/api/v1/users/:id", apihandler.ServeUser)     // Get a single user by ID
		public.Post("/api/v1/users", apihandler.CreateUser)       // Create a new user
		public.Patch("/api/v1/users/:id", apihandler.UpdateUser)  // Update user by ID
		public.Delete("/api/v1/users/:id", apihandler.DeleteUser) // Delete user by ID
	}
}

func InitProtectedRoutes(app *fiber.App) {
	apihandler := handler.ApiHandler{}
	freetier := app.Group("/api/v1")
	freetier.Use(middlewares.RequiresRealmRole("freetier"))
	{
		freetier.Get("/product-protected/:id", apihandler.ServeProduct)
	}
}
