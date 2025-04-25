#!/bin/bash

num="$1"
digits=${#num}
sum=0
temp=$num

while [ "$temp" -gt 0 ]; do
    digit=$((temp % 10))
    sum=$((sum + digit**digits))
    temp=$((temp / 10))
done

if [ "$sum" -eq "$num" ]; then
    echo "true"
else
    echo "false"
fi
