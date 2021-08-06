package graphql

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
)

func Test_GraphQL001(t *testing.T) {

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `
       {
           hello
       }
    `

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
