#!/bin/bash

# Project configuration
PROJECT_NAME="signal"

# Get current date and time
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_DIR="$(dirname "$0")"
BACKUP_FILE="$BACKUP_DIR/backup_$TIMESTAMP.sql"

# Backup the database
pg_dump -h localhost -U postgres -d "$PROJECT_NAME" >"$BACKUP_FILE"

# Compress the backup
gzip "$BACKUP_FILE"

echo "Backup created: ${BACKUP_FILE}.gz"

