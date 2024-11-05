package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/api/handler"
	"github.com/iamyassin08/prep/api/middlewares"
	"github.com/iamyassin08/prep/docs"
	"github.com/iamyassin08/prep/identity"
	"github.com/iamyassin08/prep/shagreen"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func InitPublicRoutes(app *fiber.App) {
	// apiHandler := &handler.ApiHandler{queries: db.DB}
	apihandler := handler.ApiHandler{}
	identityManager := identity.NewIdentityManager()
	registerUseCase := shagreen.NewRegistraterUseCase(identityManager)
	docs.SwaggerInfo.BasePath = "/api/v1"
	public := app.Group("/")
	{
		public.Post("/api/v1/register", handler.RegisterHandler(registerUseCase))
		public.Post("/api/v1/login", handler.LoginHandler(registerUseCase))
		public.Get("/swagger/*", fiberSwagger.WrapHandler)
		public.Get("/api/v1/healthz", handler.HealthCheck)

		public.Post("/api/v1/users/:id/thumbnail", apihandler.UploadThumbnail)
		public.Get("/api/v1/users/:id/thumbnail", apihandler.GetUserThumbnail)
		public.Get("/api/v1/users/:id/images", apihandler.GetUserImages)
		public.Post("/api/v1/users/:id/images", apihandler.UploadFile)

		public.Get("/api/v1/users", apihandler.ListUsers)
		public.Get("/api/v1/users/:id", apihandler.ServeUser)
		public.Patch("/api/v1/users/:id", apihandler.UpdateUser)

	}
}

func InitProtectedRoutes(app *fiber.App) {
	apihandler := handler.ApiHandler{}
	freetier := app.Group("/api/v1")
	freetier.Use(middlewares.RequiresRealmRole("freetier"))
	{
		freetier.Get("/user-protected/:id", apihandler.ServeUser)
	}
}
