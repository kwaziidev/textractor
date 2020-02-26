package textractor

import "math"

func mean(a []float64) float64 {
	return sum(a) / float64(len(a))
}

func sum(a []float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

func diffSqrtMean(a []float64) []float64 {
	set := make([]float64, len(a))
	meanVal := mean(a)
	var d float64
	for i, v := range a {
		d = v - meanVal
		set[i] = d * d
	}
	return set
}

func std(a []float64) float64 {
	return math.Sqrt(mean(diffSqrtMean(a)))
}
