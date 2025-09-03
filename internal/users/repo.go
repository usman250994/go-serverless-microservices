package user

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Repository interface {
	Save(ctx context.Context, user *profile) (string, error)
}

type DynamoRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoRepository(client *dynamodb.Client, tableName string) *DynamoRepository {
	return &DynamoRepository{client: client, tableName: tableName}
}

func (r *DynamoRepository) Save(ctx context.Context, user *profile) (string, error) {
	item := map[string]types.AttributeValue{
		"id":       &types.AttributeValueMemberS{Value: user.Id},
		"name":     &types.AttributeValueMemberS{Value: user.Name},
		"email":    &types.AttributeValueMemberS{Value: user.Email},
		"userType": &types.AttributeValueMemberS{Value: string(user.UserType)},
		"rating":   &types.AttributeValueMemberS{Value: fmt.Sprintf("%f", user.Rating)},
	}

	_, err := r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})
	if err != nil {
		return "_", err
	}

	return "success", nil
}
