package item

import (
	"github.com/opensearch-project/opensearch-go/v2"
)

// type Repository interface {
// 	Save(ctx context.Context, res *product) (string, error)
// 	Find(ctx context.Context, userId string) (product, error)
// }

type OpenSearchRepository struct {
	client    *opensearch.Client
	indexName string
}

func NewRepository(client *opensearch.Client, indexName string) *OpenSearchRepository {
	return &OpenSearchRepository{client: client, indexName: indexName}
}

// func (r *OpenSearchRepository) Save(ctx context.Context, user *profile) (string, error) {
// 	// abstract the OpenSearch document fields to help if repo grows for cleanliness
// 	// res-usability and to avoid repetition
// 	// e.g. if we change DB, we only change here
// 	// or if we add more fields, we only add here
// 	// instead of in every place where we save
// 	item := map[string]interface{}{
// 		"Id":       user.Id,
// 		"name":     user.Name,
// 		"email":    user.Email,
// 		"userType": string(user.UserType),
// 		"rating":   fmt.Sprintf("%f", user.Rating),
// 	}

// 	_, err := r.client.Index().Index(r.indexName).BodyJson(item).Do(ctx)
// 	if err != nil {
// 		return "_", err
// 	}

// 	return "success", nil
// }

// func (r *OpenSearchRepository) Find(ctx context.Context, userId string) (*profile, error) {
// 	// abstract this logic too
// 	// to avoid repetition and for cleanliness
// 	// if we change DB, we only change here
// 	// or if we add more fields, we only add here
// 	// instead of in every place where we fetch
// 	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
// 		TableName: aws.String(r.tableName),
// 		Key: map[string]types.AttributeValue{
// 			"Id": &types.AttributeValueMemberS{Value: userId},
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user profile
// 	if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
