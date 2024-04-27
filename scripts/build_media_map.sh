#!/bin/bash

data_dir="data"

if [[ ! -d "$data_dir" ]]; then
    echo "Error: Directory $data_dir does not exist."
    exit 1
fi

media_types=""


for media_type in "$data_dir"/*; do
    if [[ -d "$media_type" ]]; then
        media_type_name=$(basename "$media_type")
        media_types+="$media_type_name "
        
        echo "Media Type: $media_type_name"
        echo "Titles:"

        for file in "$media_type"/*.json; do
            if [[ -f "$file" ]]; then
            
                title=$(basename "$file" .json)
                echo "  - $title"
            fi
        done

        echo "" 
    fi
done

# Print all media types collected
echo "All Media Types: $media_types"
