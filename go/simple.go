package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
)

func main() {
    // Open the file specified in the first command-line argument
    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Read the numbers from the file into a slice
    var numbers []float64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        num := parseNum(scanner.Text())
        numbers = append(numbers, num)
    }

    // Compute the average
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    avg := sum / float64(len(numbers))

    // Compute the median
    sort.Float64s(numbers)
    var median float64
    if len(numbers)%2 == 0 {
        median = (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
    } else {
        median = numbers[len(numbers)/2]
    }

    // Compute the variance
    var variance float64
    for _, num := range numbers {
        variance += math.Pow(num-avg, 2)
    }
    variance /= float64(len(numbers))

    // Compute the standard deviation
    stdDev := math.Sqrt(variance)

    // Print the results
    fmt.Printf("Average: %.2f\n", avg)
    fmt.Printf("Median: %.2f\n", median)
    fmt.Printf("Variance: %.2f\n", variance)
    fmt.Printf("Standard Deviation: %.2f\n", stdDev)
}

// Parses a number from a string
func parseNum(s string) float64 {
    num, err := strconv.ParseFloat(s, 64)
    if err != nil {
        panic(err)
    }
    return num
}
