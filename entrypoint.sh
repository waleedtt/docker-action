#!/bin/sh -l

echo "Hello $1"
echo "Hello $@"

printenv

time=1
echo "cache-hit=$time" >> $GITHUB_OUTPUT