#!/bin/bash

# system-performance-monitor.sh
# This script monitors CPU, memory, disk, and network usage at regular intervals
# and logs the data into a log file.

LOG_FILE="system_performance.log"
INTERVAL=5  # Interval in seconds between checks (can be modified)

# Print the header to the log file
echo "Monitoring System Performance..." > "$LOG_FILE"
echo "Timestamp, CPU Usage (%), Memory Usage (%), Disk Usage (%), Network In (KB/s), Network Out (KB/s)" >> "$LOG_FILE"

# Function to get CPU usage percentage
get_cpu_usage() {
    echo "$(top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | awk '{print 100 - $1}')"
}

# Function to get Memory usage percentage
get_memory_usage() {
    echo "$(free | grep Mem | awk '{print $3/$2 * 100.0}')"
}

# Function to get Disk usage percentage
get_disk_usage() {
    echo "$(df / | grep / | awk '{ print $5 }' | sed 's/%//g')"
}

# Function to get Network in/out speed (KB/s)
get_network_usage() {
    local net_in=$(cat /proc/net/dev | grep eth0 | awk '{print $2}')
    local net_out=$(cat /proc/net/dev | grep eth0 | awk '{print $10}')
    
    # Convert bytes to KB and return both values
    echo "$((net_in / 1024)),$((net_out / 1024))"
}

# Main loop: Collect and log data at regular intervals
while true; do
    TIMESTAMP=$(date "+%Y-%m-%d %H:%M:%S")
    CPU=$(get_cpu_usage)
    MEMORY=$(get_memory_usage)
    DISK=$(get_disk_usage)
    NETWORK=$(get_network_usage)

    # Log the data into the log file
    echo "$TIMESTAMP, $CPU, $MEMORY, $DISK, $NETWORK" >> "$LOG_FILE"

    # Display the current data on the screen
    echo "[$TIMESTAMP] CPU: $CPU%, Memory: $MEMORY%, Disk: $DISK%, Network In/Out: $NETWORK KB/s"
    
    # Sleep for the specified interval before the next check
    sleep "$INTERVAL"
done
