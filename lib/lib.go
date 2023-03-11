package math_skills

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	check := os.Args
	if len(check) == 2 {
		fileName := check[1]

		// Get from file data
		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		// Split data
		array := strings.Fields(string(fileContent))

		// Preparation
		sort.Strings(array)
		numberOfArray := len(array)
		medianPosition := numberOfArray / 2 // for Median
		sumOfArray, dataErr := Sum(array)
		if dataErr != nil {
			log.Fatal(dataErr)
			return
		}

		// Calculation
		average := calculateAverage(sumOfArray, numberOfArray)
		variance := calculateVariance(array, average, numberOfArray)
		standardDeviation := math.Sqrt(variance)
		median := calculateMedian(array[medianPosition], array[medianPosition-1], numberOfArray)

		// Result Print
		_number := int(math.Round(average))
		number := strconv.Itoa(_number)
		fmt.Printf("Average: %s\n", number)

		_number = int(math.Round(median))
		number = strconv.Itoa(_number)
		fmt.Printf("Median: %s\n", number)

		_number = int(math.Round(variance))
		number = strconv.Itoa(_number)
		fmt.Printf("Variance: %s\n", number)

		_number = int(math.Round(standardDeviation))
		number = strconv.Itoa(_number)
		fmt.Printf("Standard Deviation: %s\n", number)

	} else {
		fmt.Println("Please enter arguments: a source.txt")
		os.Exit(1)
	}
}

func Sum(arr []string) (int, error) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return sum, nil
}
