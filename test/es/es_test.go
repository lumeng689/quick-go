package es

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

func getEsClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
		},
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	log.Println(res)

	return es
}

func TestEsConnect001(t *testing.T) {
	getEsClient()
}

func TestInsertData001(t *testing.T) {
	es := getEsClient()

	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: "1",
		Body:       strings.NewReader(`{"name": "zhangsan", "age": 18}`),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error index document", res.Status())
	} else {
		log.Println("response::", res.String())
	}
}

func TestSearchData001(t *testing.T) {
	es := getEsClient()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "zhangsan",
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error index document", res.Status())
	} else {
		log.Println("response::", res.String())
	}
}
