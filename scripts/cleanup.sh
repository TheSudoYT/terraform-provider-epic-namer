#!/bin/bash

# Define the paths for the directory and file
terraformDir=".terraform"
terraformLockFile=".terraform.lock.hcl"
terraformState="terraform.tfstate"
backupstate="terraform.tfstate.backup"

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

# Check and remove the terraform.tfstate file if it exists
if [ -f "$terraformState" ]; then
    rm -f "$terraformState"
    echo "Removed .terraform.tfstate file."
else
    echo ".terraform.tfstate file does not exist."
fi

# Check and remove the terrraform.tfstate.backup file if it exists
if [ -f "$backupstate" ]; then
    rm -f "$backupstate"
    echo "Removed terrraform.tfstate.backup  file."
else
    echo "terrraform.tfstate.backup  file does not exist."
fi
