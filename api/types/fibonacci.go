package types

import (
	"math/big"

	"github.com/graphql-go/graphql"
)

type Fibonacci struct {
	N        int      `json:"n"`
	Sequence *big.Int `json:"sequence"`
}

var FibonacciType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Fibonacci",
	Fields: graphql.Fields{
		"n": &graphql.Field{
			Type: graphql.Int,
		},
		"sequence": &graphql.Field{
			Type: graphql.String,
		},
	},
})
