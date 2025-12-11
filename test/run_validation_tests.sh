#!/bin/bash

# Validate one by one
for test in $(ls -1 validate/); do cd validate/${test}/; sfvbrr validate $(ls -1); cd ../..; done > testresult.log

# Validate recursively
# sfvbrr validate -r validate/ > testresult.log
