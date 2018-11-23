package compute

import (
	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/battle-royal/api/types"
)

// Calculate the last number of the fibonacci sequence by n iterations
func GetFibonacci() (field *graphql.Field) {
	field = &graphql.Field{
		Type: FibonacciType,
		Args: graphql.FieldConfigArgument{
			"n": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			f := Fibonacci{
				N: p.Args["n"].(int),
			}

			sequence := fibonacci(f.N)
			f.Sequence = sequence

			return f, nil
		},
	}

	return
}
