#!/bin/bash
# Network diagnostics script

echo "===== Network Diagnostics ====="

# Check if the system is online
if ping -c 1 8.8.8.8 &>/dev/null; then
    echo "Internet: Connected"
else
    echo "Internet: Disconnected"
fi

# Display the current IP address
echo "IP Address: $(hostname -I | awk '{print $1}')"

# Show active network interfaces
echo "Active Network Interfaces:"
ip -brief addr show | grep 'UP' | awk '{print $1}'

# Display DNS servers
echo "DNS Servers:"
cat /etc/resolv.conf | grep 'nameserver'

# Test connection to a website
website="google.com"
echo "Testing connection to $website..."
ping -c 3 $website

echo "================================"