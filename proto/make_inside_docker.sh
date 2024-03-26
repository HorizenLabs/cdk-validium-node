#!/bin/bash

SRC_DIR=$1 # Directory inside the docker container where source files are mounted 
git config --global --add safe.directory $SRC_DIR # necessary to safely call git commands in the source directory inside the docker container
make generate-code-from-proto
