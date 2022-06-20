package main

import (
	"encoding/json"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/require"
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
	expected := `{"data":{"hello":"world"}}`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
		t.FailNow()
	}
	rJSON, _ := json.Marshal(r)
	require.JSONEq(t, expected, string(rJSON), "hello world comparison failed: %s", "formatted")
}
