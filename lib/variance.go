package math_skills

import (
	"fmt"
	"math"
	"strconv"
)

func calculateVariance(arr []string, avg float64, n int) float64 {
	v := 0.0
	for i := 0; i < len(arr); i++ {
		num, err := strconv.ParseFloat(arr[i], 64)
		if err != nil {
			fmt.Println(arr[i] + "is mot number")
			return 0
		}
		v += math.Pow(num-avg, 2)
	}
	return v / float64(n)
}
