#!/usr/bin/env bash

main() {

	result=""

	if [ -z "$1" ]; then
		echo "Please provide a number as an argument."
		exit 1
	fi

	number="$1"

	if ! [[ "$number" =~ ^[0-9]+$ ]]; then
		echo "Invalid input, please enter a valid number"
		exit 1
	fi

	if (( number % 3 == 0)); then
		result+="Pling"
	fi

	if (( number % 5 == 0)); then
		result+="Plang"
	fi

	if (( number % 7 == 0 )); then
		result+="Plong"
	fi

	if [ -z "$result" ]; then
		result="$number"
	fi

	echo $result
}

main "$@"
