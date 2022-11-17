#!/bin/sh -l

echo "Hello $1"
name='INPUT_WHO-TO-GREET'
echo "INPUT ${!name}"


time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT