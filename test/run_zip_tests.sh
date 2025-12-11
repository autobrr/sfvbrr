#!/bin/bash

# Check zip one by one
for test in $(ls -1 zip/); do sfvbrr zip zip/${test}; done 2>&1> testresult.log

# Check zip recursively
# sfvbrr zip -r zip/ > testresult.log
