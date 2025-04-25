#!/usr/bin/env bash

input="$*"
input="${input//-/ }"
input=$(echo "$input" | tr -d '[:punct:]')

acronym=""
for word in $input; do
	first_letter="${word:0:1}"
	acronym+=$(echo "$first_letter" | tr '[:lower:]' '[:upper:]')
done

echo "$acronym"
