#!/bin/bash
# Backup Script

echo "===== Backup Script ====="

# Define the source and destination directories
SOURCE_DIR="/home/user/documents"
DEST_DIR="/home/user/backup"

# Create a timestamp for the backup folder
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_DIR="$DEST_DIR/backup_$TIMESTAMP"

# Create the destination directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

# Perform the backup using rsync
echo "Backing up files from $SOURCE_DIR to $BACKUP_DIR..."
rsync -av --progress "$SOURCE_DIR" "$BACKUP_DIR"

# Verify backup
if [ $? -eq 0 ]; then
    echo "Backup completed successfully!"
else
    echo "Backup failed!"
fi

echo "==================================="
