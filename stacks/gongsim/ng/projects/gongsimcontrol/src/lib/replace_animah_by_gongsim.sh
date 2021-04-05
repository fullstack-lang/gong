#!/bin/bash

for file in $(find . -name 'animah*')
do
  mv $file $(echo "$file" | sed 's|animah|gongsim|g')
  # echo "$file" | sed 's|animah|gongsim|g'
done
