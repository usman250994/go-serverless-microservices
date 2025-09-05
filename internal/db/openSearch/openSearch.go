package openSearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

func NewClient(cfg aws.Config, endpoint string, tbl string) (*opensearch.Client, error) {
	signer, err := requestsigner.NewSignerWithService(cfg, "aoss")
	if err != nil {
		return nil, fmt.Errorf("signer error: %w", err)
	}

	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{endpoint},
		Signer:    signer,
	})
	if err != nil {
		return nil, fmt.Errorf("client creation error: %w", err)
	}

	// Only create index if it doesn't exist
	indexName := tbl
	settings := map[string]interface{}{
		"settings": map[string]interface{}{
			"index": map[string]interface{}{
				"number_of_shards":   1,
				"number_of_replicas": 0,
			},
		},
	}
	settingsJSON, _ := json.Marshal(settings)

	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  bytes.NewReader(settingsJSON),
	}
	resp, err := req.Do(context.Background(), client)

	if err != nil {
		// If index already exists, ignore error
		if resp != nil && resp.StatusCode == 400 {
			log.Println("Index already exists, skipping creation.")
		} else {
			return nil, fmt.Errorf("failed to create index: %w", err)
		}
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	return client, nil
}
