package compute

import (
	"math/big"
)

func fibonacci(n int) (f *big.Int) {
	f = big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		f.Add(f, b)

		f, b = b, f
	}

	return
}
