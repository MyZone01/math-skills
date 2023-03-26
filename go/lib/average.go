package math_skills

import "math"

func calculateAverage(sum, n int) float64 {
	return math.Round(float64(sum) / float64(n))
}
