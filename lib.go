package a

import (
	"math"

	"github.com/cybersamx/c"
)

func StdDev(sample []float64) float64 {
	sum := 0.0
	mean := c.Mean(sample)

	for _, v := range sample {
		sum += math.Pow(v-mean, 2)
	}

	stddev := math.Sqrt(sum / c.Count(sample))

	return stddev
}
