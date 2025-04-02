#!/usr/bin/env bash

main() {

	if [ $# -eq 1 ]; then
		echo "Hello, $1"
	elif [ $# -eq 0 ]; then
		echo "Usage: error_handling.sh <person>"
		exit 1
	else
		echo "Usage: error_handling.sh <person>"
		exit 1
	fi
}

main "$@" 
