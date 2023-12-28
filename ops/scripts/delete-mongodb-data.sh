#!/bin/bash

directory="ops/docker/data"
file_to_keep=".gitignore"

if [ -d "$directory" ]; then
  if [ -e "$directory/$file_to_keep" ]; then
    for file in "$directory"/*; do
      if [ -f "$file" ] && [ "$file" != "$directory/$file_to_keep" ]; then
        rm -f "$file"
      fi
    done

    for subdir in "$directory"/*/; do
      if [ -d "$subdir" ]; then
        rm -rf "$subdir"
      fi
    done
  else
    echo "The file '$file_to_keep' does not exist in the directory."
  fi
else
  echo "The directory '$directory' does not exist."
fi
