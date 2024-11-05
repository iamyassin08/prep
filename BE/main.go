package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iamyassin08/prep/api/middlewares"
	"github.com/iamyassin08/prep/api/routes"
	"github.com/iamyassin08/prep/db"
	"github.com/iamyassin08/prep/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"

	// _ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var tracer = otel.Tracer("prep")

func initTracer() *sdktrace.TracerProvider {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceFirstNameKey.String("prep"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	ctx := context.Background()
	// tp := initTracer()
	// defer func() {
	// 	if err := tp.Shutdown(ctx); err != nil {
	// 		log.Printf("Error shutting down tracer provider")
	// 	}
	// }()
	stripeKey, _ := os.LookupEnv("PREP_DB_HOST")
	params := &stripe.ChargeParams{}
	sc := &client.API{}
	sc.Init(stripeKey, nil)
	sc.Charges.Get("ch_3Ln3j02eZvKYlo2C0d5IZWuG", params)
	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
		AppName:           "Prep",
		ServerHeader:      "Fiber",
	})
	app.Use(otelfiber.Middleware())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	dB, err := pgxpool.New(ctx, GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	db.DB = db.New(dB)
	storage.MINIO_CLIENT, err = storage.ConnectToMinioClient()
	if err != nil {
		fmt.Println(err)
	}
	defer dB.Close()
	// bkex, _ := storage.MINIO_CLIENT.BucketExists(context.Background(), "prep-bucket")
	// _ = storage.MINIO_CLIENT.MakeBucket(ctx, "prep", minio.MakeBucketOptions{Region: "us-central"})
	fmt.Println(storage.MINIO_CLIENT.ListBuckets(ctx))
	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, routes.InitProtectedRoutes)
	log.Fatal(app.Listen(":8080"))
}

func GetConnectionString() string {
	dbHost, _ := os.LookupEnv("PREP_DB_HOST")
	dbFirstName, _ := os.LookupEnv("PREP_DB_NAME")
	dbPass, _ := os.LookupEnv("PREP_DB_PASS")
	dbUser, _ := os.LookupEnv("PREP_DB_USER")
	dbSsl, _ := os.LookupEnv("PREP_DB_SSL_MODE")
	if len(dbSsl) <= 2 {
		dbSsl = "disable"
	}
	dbPort, _ := os.LookupEnv("PREP_DB_PORT")
	if len(dbPort) < 1 {
		dbPort = "5432"
	}
	var db_connection_string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbFirstName)
	return db_connection_string
}
