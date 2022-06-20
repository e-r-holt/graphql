package main

import (
	"encoding/json"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestStartGraphQL(t *testing.T) { // Query

	schema, err := startGraphql()
	if err != nil {
		t.Errorf("failed to create new schema, error: %v", err)
	}

	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	t.Logf("%s \n", rJSON) // {"data":{"hello":"world"}}
}