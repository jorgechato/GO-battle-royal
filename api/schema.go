package api

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/battle-royal/api/compute"
)

func buildSchema() (schema graphql.Schema) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"fibonacci": compute.GetFibonacci(),
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: graphql.Fields{},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return
}
