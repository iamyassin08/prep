package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/api/handler"
	"github.com/iamyassin08/prep/api/middlewares"
	"github.com/iamyassin08/prep/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func InitPublicRoutes(app *fiber.App) {
	// apiHandler := &handler.ApiHandler{queries: db.DB}
	apihandler := handler.ApiHandler{}
	docs.SwaggerInfo.BasePath = "/api/v1"
	public := app.Group("/")
	{
		// Public routes: registration, login, and health check

		public.Get("/swagger/*", fiberSwagger.WrapHandler)

		// User-related routes (as per the request)
		public.Get("/api/v1/users", apihandler.ListUsers)
		public.Get("/api/v1/users/:id", apihandler.ServeUser)
		public.Post("/api/v1/users", apihandler.CreateUser)
		public.Patch("/api/v1/users/:id", apihandler.UpdateUser)
		public.Delete("/api/v1/users/:id", apihandler.DeleteUser)
	}
}

func InitProtectedRoutes(app *fiber.App) {
	freetier := app.Group("/api/v1")
	freetier.Use(middlewares.RequiresRealmRole("freetier"))
	{
		freetier.Get("/user/:id", apihandler.ServeUser)
	}
}
