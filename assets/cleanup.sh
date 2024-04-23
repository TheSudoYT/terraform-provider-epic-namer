#!/bin/bash

# Define the paths for the directory and file
terraformDir=".terraform"
terraformLockFile=".terraform.lock.hcl"

# Check and remove the .terraform directory if it exists
if [ -d "$terraformDir" ]; then
    rm -rf "$terraformDir"
    echo "Removed .terraform directory."
else
    echo ".terraform directory does not exist."
fi

# Check and remove the .terraform.lock.hcl file if it exists
if [ -f "$terraformLockFile" ]; then
    rm -f "$terraformLockFile"
    echo "Removed .terraform.lock.hcl file."
else
    echo ".terraform.lock.hcl file does not exist."
fi
