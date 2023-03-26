import sys
import numpy as np

# Check if a file path was passed as an argument
if len(sys.argv) < 2:
    print("Usage: python statistics.py <file_path>")
    sys.exit()

# Read data from the file
data = np.loadtxt(sys.argv[1])

# Calculate statistics
mean = np.mean(data)
median = np.median(data)
variance = np.var(data)
std_dev = np.std(data)

# Print the results
print("Mean:", mean)
print("Median:", median)
print("Variance:", variance)
print("Standard deviation:", std_dev)
