package item

import (
	"context"
	"fmt"
	"strings"

	"bytes"
	"encoding/json"

	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

type Repository interface {
	Save(ctx context.Context, req *product) (string, error)
	getNearestTen(ctx context.Context, q *ProductQuery) ([]product, error)
}

type OpenSearchRepository struct {
	client    *opensearch.Client
	indexName string
}

func NewRepository(client *opensearch.Client, indexName string) *OpenSearchRepository {
	return &OpenSearchRepository{client: client, indexName: indexName}
}

func (r *OpenSearchRepository) Save(ctx context.Context, prd *product) (string, error) {
	// abstract the OpenSearch document fields to help if repo grows for cleanliness
	// res-usability and to avoid repetition
	// e.g. if we change DB, we only change here
	// or if we add more fields, we only add here
	// instead of in every place where we save

	// marshal prd to JSON and pass as io.Reader
	data, err := json.Marshal(prd)
	if err != nil {
		return "", err
	}
	_, err = r.client.Index(r.indexName, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (r *OpenSearchRepository) getNearestTen(ctx context.Context, q *ProductQuery) ([]product, error) {
	// abstract this logic too
	// to avoid repetition and for cleanliness
	// if we change DB, we only change here
	// or if we add more fields, we only add here
	// instead of in every place where we fetch

	// query to find nearest 10 products from given lat/lng and match name/details if provided

	// later abstract this query into opensearch package to avoid repetition and future changes
	// for now, keep it here for simplicity
	// as we have only one query
	query := `{
		"size": 10,
		"query": {
			"bool": {
				"must": [
					{
						"geo_distance": {
							"distance": "200km",
							"location": {
								"lat": ` + fmt.Sprintf("%f", q.Lat) + `,
								"lon": ` + fmt.Sprintf("%f", q.Lng) + `
							}
						}
					}
				]
			}
		}
	}`

	// use this query to search
	req := opensearchapi.SearchRequest{
		Index: []string{r.indexName},
		Body:  strings.NewReader(query),
	}

	// return array of products or err
	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var products []product
	if err := json.NewDecoder(res.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}
