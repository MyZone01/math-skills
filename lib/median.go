package math_skills

import (
	"fmt"
	"strconv"
)

func calculateMedian(a, b string, n int) float64 {
	c, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(a + " is not number")
		return 0
	}
	if n%2 != 0 {
		return float64(c)
	}

	d, err2 := strconv.ParseFloat(b, 64)

	if err2 != nil {
		fmt.Println(b + " is not number")
		return 0
	}
	return (c + d) / float64(2)
}
