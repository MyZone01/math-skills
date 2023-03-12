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
	if len(os.Args) == 2 {
		fileName := os.Args[1]

		// Get from file data
		data := getDataFromFile(fileName)

		// Sort data for median
		sort.Strings(data)
		numberOfData := len(data)
		medianPosition := numberOfData / 2 // for Median
		sumOfArray, dataErr := Sum(data)
		if dataErr != nil {
			log.Fatal(dataErr)
			return
		}

		// Calculation
		average := calculateAverage(sumOfArray, numberOfData)
		variance := calculateVariance(data, average, numberOfData)
		standardDeviation := math.Sqrt(variance)
		median := calculateMedian(data[medianPosition], data[medianPosition-1], numberOfData)

		// Result Print
		number := strconv.Itoa(int(math.Round(average)))
		fmt.Printf("Average: %s\n", number)
		number = strconv.Itoa(int(math.Round(median)))
		fmt.Printf("Median: %s\n", number)
		number = strconv.Itoa(int(math.Round(variance)))
		fmt.Printf("Variance: %s\n", number)
		number = strconv.Itoa(int(math.Round(standardDeviation)))
		fmt.Printf("Standard Deviation: %s\n", number)
	}
}

func getDataFromFile(fileName string) []string {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Fields(string(fileContent))
	return data
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
