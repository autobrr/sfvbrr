#!/bin/bash

# Check sfv one by one
for test in $(ls -1 sfv/); do sfvbrr sfv sfv/${test}; done 2>&1> testresult.log

# Check sfv recursively
# sfvbrr sfv -r sfv/ > testresult.log
