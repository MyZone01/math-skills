import 'dart:io';
import 'dart:math';

void main(List<String> arguments) {
  // Check if a file path was passed as an argument
  if (arguments.length < 1) {
    print('Usage: dart statistics.dart <file_path>');
    exit(1);
  }

  // Read data from the file
  final data = File(arguments[0]).readAsLinesSync().map(double.parse).toList();

  // Calculate statistics
  final mean = data.reduce((a, b) => a + b) / data.length;
  final sortedData = [...data]..sort();
  final median = sortedData.length % 2 == 0
      ? (sortedData[sortedData.length ~/ 2 - 1] +
          sortedData[sortedData.length ~/ 2]) /
          2
      : sortedData[sortedData.length ~/ 2];
  final variance = data
      .map((x) => pow(x - mean, 2))
      .reduce((a, b) => a + b) /
      data.length;
  final stdDev = sqrt(variance);

  // Print the results
  print('Mean: $mean');
  print('Median: $median');
  print('Variance: $variance');
  print('Standard deviation: $stdDev');
}
