package user

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Repository interface {
	Save(ctx context.Context, user *profile) (string, error)
	FindById(ctx context.Context, userId string) (*profile, error)
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
		"Id":       &types.AttributeValueMemberS{Value: user.Id},
		"name":     &types.AttributeValueMemberS{Value: user.Name},
		"email":    &types.AttributeValueMemberS{Value: user.Email},
		"userType": &types.AttributeValueMemberS{Value: string(user.UserType)},
		"rating":   &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", user.Rating)},
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

func (r *DynamoRepository) FindById(ctx context.Context, userId string) (*profile, error) {
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: userId},
		},
	})
	if err != nil {
		return nil, err
	}

	var user profile
	if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
