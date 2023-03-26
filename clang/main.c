#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(int argc, char *argv[]) {
    // Check if a file path was passed as an argument
    if (argc < 2) {
        printf("Usage: statistics <file_path>\n");
        return 1;
    }
    
    // Open the file
    FILE *fp = fopen(argv[1], "r");
    if (fp == NULL) {
        printf("Error: could not open file\n");
        return 1;
    }
    
    // Read data from the file
    int n = 0;
    double sum = 0;
    double *data = malloc(sizeof(double));
    double x;
    while (fscanf(fp, "%lf", &x) == 1) {
        n++;
        sum += x;
        data = realloc(data, n * sizeof(double));
        data[n-1] = x;
    }
    
    // Calculate statistics
    double average = sum / n;
    double median;
    double variance = 0;
    double std_dev;
    int i, j;
    
    // Calculate median
    for (i = 0; i < n-1; i++) {
        for (j = i+1; j < n; j++) {
            if (data[j] < data[i]) {
                double temp = data[i];
                data[i] = data[j];
                data[j] = temp;
            }
        }
    }
    if (n % 2 == 0) {
        median = (data[n/2-1] + data[n/2]) / 2;
    } else {
        median = data[n/2];
    }
    
    // Calculate variance and standard deviation
    for (i = 0; i < n; i++) {
        variance += pow(data[i] - average, 2);
    }
    variance /= n;
    std_dev = sqrt(variance);
    
    // Print the results
    printf("Average: %d\n", (int) round(average));
    printf("Median: %d\n", (int) round(median));
    printf("Variance: %d\n", (int) round(variance));
    printf("Standard deviation: %d\n", (int) round(std_dev));
    
    // Clean up
    free(data);
    fclose(fp);
    
    return 0;
}
