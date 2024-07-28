#!/bin/bash

# File to store the patterns
PATTERNSFILE="patterns.txt"

# Pattern lengths
PATTERN_LENGTHS=(3 10 50 100 200 500 1000 2000 5000 10000)

# Function to generate a random pattern of given length
generate_pattern() {
    local length=$1
    tr -dc 'a-z' < /dev/urandom | head -c $length
}

# Create or clear the patterns file
> $PATTERNSFILE

# Generate patterns and write to the file
for length in "${PATTERN_LENGTHS[@]}"; do
    generate_pattern $length >> $PATTERNSFILE
    echo "" >> $PATTERNSFILE  # Add a newline after each pattern
done

echo "Patterns generated in $PATTERNSFILE"
