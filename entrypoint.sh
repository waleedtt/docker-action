#!/bin/sh -l

echo "Hello $1"
echo "INPUT ${INPUT_WHO-TO-GREET}"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT