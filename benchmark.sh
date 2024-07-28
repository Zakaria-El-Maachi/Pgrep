#!/bin/bash

# File containing the patterns
PATTERNSFILE="patterns.txt"
# Big text file to search within
TEXTFILE="bigfile.txt"
# Compiled Go binary
GO_BINARY="pgrep"

# Function to calculate mean
calculate_mean() {
    awk '{sum+=$1} END {print sum/NR}'
}

# Function to calculate std deviation
calculate_std_dev() {
    awk '{sum+=$1; sumsq+=$1*$1} END {print sqrt(sumsq/NR - (sum/NR)^2)}'
}

# Output file
OUTPUTFILE="benchmark_results.txt"
echo "Pattern Length,Program,Mean Time (s),Std Dev Time (s)" > $OUTPUTFILE

# Read patterns from the file
patterns=()
while IFS= read -r line; do
    patterns+=("$line")
done < "$PATTERNSFILE"

# Run tests and collect timing data
for pattern in "${patterns[@]}"; do
    pattern_length=${#pattern}

    # Test grep -oba
    echo "Testing grep -oba with pattern length $pattern_length"
    grep_times=()
    for ((i=0; i<10; i++)); do
        grep_time=$( { time grep -oba "$pattern" $TEXTFILE > /dev/null; } 2>&1 | grep real | awk '{print $2}' | sed 's/.*m\([0-9.]*\)s/\1/')
        grep_times+=($grep_time)
    done
    grep_mean=$(printf "%s\n" "${grep_times[@]}" | calculate_mean)
    grep_std_dev=$(printf "%s\n" "${grep_times[@]}" | calculate_std_dev)
    echo "Pattern length $pattern_length: grep - Mean: $grep_mean s, Std Dev: $grep_std_dev s"
    echo "$pattern_length,grep -oba,$grep_mean,$grep_std_dev" >> $OUTPUTFILE

    # Test Go program without -p
    echo "Testing Go program without -p with pattern length $pattern_length"
    go_times=()
    for ((i=0; i<10; i++)); do
        go_time=$( { time ./$GO_BINARY "$pattern" $TEXTFILE > /dev/null; } 2>&1 | grep real | awk '{print $2}' | sed 's/.*m\([0-9.]*\)s/\1/')
        go_times+=($go_time)
    done
    go_mean=$(printf "%s\n" "${go_times[@]}" | calculate_mean)
    go_std_dev=$(printf "%s\n" "${go_times[@]}" | calculate_std_dev)
    echo "Pattern length $pattern_length: Go program - Mean: $go_mean s, Std Dev: $go_std_dev s"
    echo "$pattern_length,Go program,$go_mean,$go_std_dev" >> $OUTPUTFILE

    # Test Go program with -p
    echo "Testing Go program with -p with pattern length $pattern_length"
    go_p_times=()
    for ((i=0; i<10; i++)); do
        go_p_time=$( { time ./$GO_BINARY -p "$pattern" $TEXTFILE > /dev/null; } 2>&1 | grep real | awk '{print $2}' | sed 's/.*m\([0-9.]*\)s/\1/')
        go_p_times+=($go_p_time)
    done
    go_p_mean=$(printf "%s\n" "${go_p_times[@]}" | calculate_mean)
    go_p_std_dev=$(printf "%s\n" "${go_p_times[@]}" | calculate_std_dev)
    echo "Pattern length $pattern_length: Go program -p - Mean: $go_p_mean s, Std Dev: $go_p_std_dev s"
    echo "$pattern_length,Go program -p,$go_p_mean,$go_p_std_dev" >> $OUTPUTFILE
done

echo "Benchmarking completed. Results are stored in $OUTPUTFILE."
