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
		public.Post("/api/v1/register", handler.RegisterHandler(registerUseCase))
		public.Post("/api/v1/login", handler.LoginHandler(registerUseCase))
		public.Get("/swagger/*", fiberSwagger.WrapHandler)
		public.Get("/api/v1/healthz", handler.HealthCheck)

		public.Post("/api/v1/products/:id/thumbnail", apihandler.UploadThumbnail)
		public.Get("/api/v1/products/:id/thumbnail", apihandler.GetProductThumbnail)
		public.Get("/api/v1/products/:id/images", apihandler.GetProductImages)
		public.Post("/api/v1/products/:id/images", apihandler.UploadFile)

		public.Get("/api/v1/products", apihandler.ListProducts)
		public.Get("/api/v1/products/:id", apihandler.ServeProduct)
		public.Patch("/api/v1/products/:id", apihandler.UpdateProduct)

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
