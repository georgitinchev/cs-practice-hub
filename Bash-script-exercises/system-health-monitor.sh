#!/bin/bash

# system-health-monitor.sh
# This script monitors system health including disk, memory, CPU usage,
# and checks if processes exceed a set threshold for CPU or memory consumption.
# It sends email alerts if thresholds are exceeded.

# Email settings for alerting (modify with your email settings)
ALERT_EMAIL="youremail@example.com"
EMAIL_SUBJECT="System Health Alert"

# Thresholds for critical conditions
CPU_THRESHOLD=90  # CPU usage percentage threshold
MEMORY_THRESHOLD=85  # Memory usage percentage threshold
DISK_THRESHOLD=90  # Disk usage percentage threshold
PROCESS_THRESHOLD=500000  # Process memory usage (KB) for alert

# Interval in seconds to check system health
INTERVAL=60

# Log file to store system health logs
LOG_FILE="system_health.log"

# Function to send email alerts
send_alert() {
    local message="$1"
    echo "$message" | mail -s "$EMAIL_SUBJECT" "$ALERT_EMAIL"
}

# Function to check disk usage
check_disk_usage() {
    local disk_usage=$(df / | grep / | awk '{ print $5 }' | sed 's/%//g')
    if [ "$disk_usage" -gt "$DISK_THRESHOLD" ]; then
        local message="ALERT: Disk usage is at $disk_usage%, which exceeds the threshold of $DISK_THRESHOLD%."
        echo "$message"
        send_alert "$message"
    fi
}

# Function to check CPU usage
check_cpu_usage() {
    local cpu_usage=$(top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | awk '{print 100 - $1}')
    if [ "$(echo "$cpu_usage > $CPU_THRESHOLD" | bc)" -eq 1 ]; then
        local message="ALERT: CPU usage is at $cpu_usage%, which exceeds the threshold of $CPU_THRESHOLD%."
        echo "$message"
        send_alert "$message"
    fi
}

# Function to check memory usage
check_memory_usage() {
    local memory_usage=$(free | grep Mem | awk '{print $3/$2 * 100.0}')
    if [ "$(echo "$memory_usage > $MEMORY_THRESHOLD" | bc)" -eq 1 ]; then
        local message="ALERT: Memory usage is at $memory_usage%, which exceeds the threshold of $MEMORY_THRESHOLD%."
        echo "$message"
        send_alert "$message"
    fi
}

# Function to check large running processes
check_large_processes() {
    local processes=$(ps aux --sort=-%mem | awk '{if($3>10) print $0}' | head -n 5)
    if [ ! -z "$processes" ]; then
        local message="ALERT: High resource-consuming processes detected:\n$processes"
        echo "$message"
        send_alert "$message"
    fi
}

# Main loop: Collect and log data at regular intervals
while true; do
    # Timestamp for logging
    TIMESTAMP=$(date "+%Y-%m-%d %H:%M:%S")

    # Perform system health checks
    check_disk_usage
    check_cpu_usage
    check_memory_usage
    check_large_processes

    # Log system health status to file
    echo "$TIMESTAMP - Health check completed." >> "$LOG_FILE"

    # Display system status on the screen
    echo "[$TIMESTAMP] System Health Check Completed. Logs saved to $LOG_FILE."

    # Sleep for the specified interval before the next check
    sleep "$INTERVAL"
done
