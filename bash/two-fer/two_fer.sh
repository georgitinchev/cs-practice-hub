#!/usr/bin/env bash

print_message() {
    if [[ -n "$1" ]]; then
        echo "One for $1, one for me."
    else
        echo "One for you, one for me."
    fi
}

print_message "$@"
