package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	user "github.com/usman250994/cloudyGo/internal/users"
)

func main() {

	ctx := context.Background()

	// 1️⃣ Load AWS config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	// 2️⃣ Initialize DynamoDB client
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// 3️⃣ Create repository
	repo := user.NewDynamoRepository(dynamoClient, os.Getenv("USERS_TABLE"))

	// 4️⃣ Create service
	service := user.NewService(repo)

	// 5️⃣ Create handler
	handler := user.NewHandler(service)

	// 6️⃣ Setup routes (simple net/http example)

	// Chi router
	r := chi.NewRouter()

	// Adding global middleware
	r.Use(middleware.Logger)

	// Register domain routes
	user.RegisterRoutes(r, handler)
	// Lambda adapter
	adapter := chiadapter.New(r)

	lambda.Start(adapter.ProxyWithContext)
}
