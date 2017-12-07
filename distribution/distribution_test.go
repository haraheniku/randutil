package distribution

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const N = 100000

func TestBeta(t *testing.T) {
	alpha, beta := 100., 100.

	var mu, sigma2 float64
	for i := 0; i < N; i++ {
		x := Beta(alpha, beta)
		mu += x
		sigma2 += x * x
	}
	mu = mu / N
	sigma2 = sigma2/N - mu*mu

	assert.InDelta(t, alpha/(alpha+beta), mu, 0.0002)
	assert.InDelta(t, alpha*beta/(math.Pow(alpha*beta, 2)*(alpha+beta+1)), sigma2, 0.002)
}
