#!/bin/bash

# Tells Golang where to put the executable
export GOBIN=$HOME/Documents/FCC/Assignment2/RSA/bin

# Go into src Directory and compile
cd src/RSA
echo "install" | xargs go
