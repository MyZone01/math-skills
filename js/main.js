const fs = require('fs');

// Check if a file path was passed as an argument
if (process.argv.length < 3) {
    console.log('Usage: node statistics.js <file_path>');
    process.exit(1);
}

// Read data from the file
const data = fs.readFileSync(process.argv[2], 'utf8').split('\n').map(Number).filter(x => !isNaN(x));

// Calculate statistics
const mean = data.reduce((sum, x) => sum + x, 0) / data.length;
const sortedData = data.sort((a, b) => a - b);
const median = sortedData.length % 2 === 0 ? (sortedData[sortedData.length / 2 - 1] + sortedData[sortedData.length / 2]) / 2 : sortedData[Math.floor(sortedData.length / 2)];
const variance = data.reduce((sum, x) => sum + (x - mean) ** 2, 0) / data.length;
const stdDev = Math.sqrt(variance);

// Print the results
console.log('Mean:', Math.round(mean));
console.log('Median:', Math.round(median));
console.log('Variance:', Math.round(variance));
console.log('Standard deviation:', Math.round(stdDev));
