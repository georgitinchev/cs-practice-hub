#!/bin/bash

# This script monitors disk usage and sends a warning if usage exceeds a threshold.

# Set the threshold (in percentage)
THRESHOLD=80

# Get the current disk usage percentage for the root partition
disk_usage=$(df / | grep / | awk '{ print $5 }' | sed 's/%//g')

# Check if the disk usage exceeds the threshold
if [ "$disk_usage" -gt "$THRESHOLD" ]; then
    echo "Warning: Disk usage is above threshold. Current usage is ${disk_usage}%."
else
    echo "Disk usage is normal. Current usage is ${disk_usage}%. Keep it up!"
fi
