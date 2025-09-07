package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	awsOs "github.com/usman250994/cloudyGo/internal/db/openSearch"
	items "github.com/usman250994/cloudyGo/internal/items"
)

func main() {

	// Load .env only in local/dev (not in Lambda)
	if os.Getenv("LAMBDA_TASK_ROOT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found or error loading .env (this is fine in production):", err)
		}
	}

	ctx := context.Background()

	// 1️⃣ Load AWS config (uses env vars for credentials/region)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	// 2️⃣ Initialize OpenSearch client
	tableName := os.Getenv("OPEN_SEARCH_INDEX")
	db, err := awsOs.NewClient(cfg, os.Getenv("OPEN_SEARCH_END_POINT"), tableName)
	if err != nil {
		log.Fatalf("failed to create OpenSearch client: %v", err)
	}

	repo := items.NewRepository(db, tableName)

	// 4️⃣ Create service
	service := items.NewService(repo)

	// 5️⃣ Create handler
	handler := items.NewHandler(service)

	// 6️⃣ Setup routes (simple net/http example)

	// Chi router
	r := chi.NewRouter()

	// Adding global middleware
	r.Use(middleware.Logger)

	// Register domain routes
	items.RegisterRoutes(r, handler)

	if os.Getenv("LOCAL") == "true" {
		log.Println("Running in local HTTP mode on :8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	} else {
		// Lambda adapter
		adapter := chiadapter.New(r)
		lambda.Start(adapter.ProxyWithContext)
	}
}
