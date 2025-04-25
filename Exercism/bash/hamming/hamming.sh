#!/bin/bash

# Check if number of argumetns is correct
if [[ $# -ne 2 ]]; then
    echo "Usage: hamming.sh <string1> <string2>"
    exit 1
fi

strand1=$1
strand2=$2

# Check lengths are equal
if [[ ${#strand1} -ne ${#strand2} ]]; then
	echo "strands must be of equal length"
	exit 1
fi

# initialize distance cntr
distance=0

# comparison 
for ((i=0;i<${#strand1};i++)); do
    if [[ "${strand1:i:1}" != "${strand2:i:1}" ]]; then
	    ((distance++))
    fi
done

echo $distance
