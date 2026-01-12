#!/usr/bin/env bash
# Database sync script - syncs newer database between local and production

set -e

SERVER="admin@50.116.26.167"
REMOTE_DB="/home/admin/portfolio/portfolio.db"
LOCAL_DB="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/portfolio.db"
BACKUP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/.db-backups"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "========================================="
echo "Database Sync Script"
echo "========================================="
echo ""

# Check if local database exists
if [ ! -f "$LOCAL_DB" ]; then
    echo -e "${RED}❌ Local database not found at: $LOCAL_DB${NC}"
    exit 1
fi

# Get remote database modification time
echo -e "${BLUE}📊 Checking database timestamps...${NC}"
REMOTE_TIMESTAMP=$(ssh "$SERVER" "stat -c %Y $REMOTE_DB 2>/dev/null || echo 0")

if [ "$REMOTE_TIMESTAMP" = "0" ]; then
    echo -e "${RED}❌ Remote database not found or cannot access${NC}"
    exit 1
fi

# Get local database modification time
LOCAL_TIMESTAMP=$(stat -f %m "$LOCAL_DB" 2>/dev/null || stat -c %Y "$LOCAL_DB" 2>/dev/null)

# Convert timestamps to human-readable format
if [[ "$OSTYPE" == "darwin"* ]]; then
    REMOTE_DATE=$(date -r "$REMOTE_TIMESTAMP" "+%Y-%m-%d %H:%M:%S")
    LOCAL_DATE=$(date -r "$LOCAL_TIMESTAMP" "+%Y-%m-%d %H:%M:%S")
else
    REMOTE_DATE=$(date -d "@$REMOTE_TIMESTAMP" "+%Y-%m-%d %H:%M:%S")
    LOCAL_DATE=$(date -d "@$LOCAL_TIMESTAMP" "+%Y-%m-%d %H:%M:%S")
fi

echo ""
echo "Local database:  $LOCAL_DATE"
echo "Remote database: $REMOTE_DATE"
echo ""

# Determine sync direction
if [ "$LOCAL_TIMESTAMP" -gt "$REMOTE_TIMESTAMP" ]; then
    SYNC_DIRECTION="push"
    NEWER="local"
    OLDER="remote"
    TIME_DIFF=$((LOCAL_TIMESTAMP - REMOTE_TIMESTAMP))
elif [ "$REMOTE_TIMESTAMP" -gt "$LOCAL_TIMESTAMP" ]; then
    SYNC_DIRECTION="pull"
    NEWER="remote"
    OLDER="local"
    TIME_DIFF=$((REMOTE_TIMESTAMP - LOCAL_TIMESTAMP))
else
    echo -e "${GREEN}✓ Databases are already in sync (same modification time)${NC}"
    exit 0
fi

# Calculate time difference in human-readable format
HOURS=$((TIME_DIFF / 3600))
MINUTES=$(((TIME_DIFF % 3600) / 60))
SECONDS=$((TIME_DIFF % 60))

echo -e "${YELLOW}⚠ The $NEWER database is newer by: ${HOURS}h ${MINUTES}m ${SECONDS}s${NC}"
echo ""

if [ "$SYNC_DIRECTION" = "push" ]; then
    echo -e "${BLUE}➜ Syncing: Push local database to remote (local → remote)${NC}"
else
    echo -e "${BLUE}➜ Syncing: Pull remote database to local (remote → local)${NC}"
fi
echo ""

# Create backup directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

# Perform sync
if [ "$SYNC_DIRECTION" = "push" ]; then
    # Backup remote database
    echo -e "${BLUE}📦 Backing up remote database...${NC}"
    BACKUP_NAME="remote-backup-$(date +%Y%m%d-%H%M%S).db"
    ssh "$SERVER" "mkdir -p /home/admin/portfolio/.db-backups && cp $REMOTE_DB /home/admin/portfolio/.db-backups/$BACKUP_NAME"
    echo -e "${GREEN}✓ Remote backup created: $BACKUP_NAME${NC}"

    # Push local to remote
    echo -e "${BLUE}⬆ Pushing local database to remote...${NC}"
    rsync -avz "$LOCAL_DB" "$SERVER:$REMOTE_DB"
    echo -e "${GREEN}✓ Database pushed successfully${NC}"

    # Restart remote service to pick up changes
    echo -e "${BLUE}🔄 Restarting remote service...${NC}"
    ssh "$SERVER" "sudo systemctl restart portfolio"
    echo -e "${GREEN}✓ Service restarted${NC}"

else
    # Backup local database
    echo -e "${BLUE}📦 Backing up local database...${NC}"
    BACKUP_NAME="local-backup-$(date +%Y%m%d-%H%M%S).db"
    cp "$LOCAL_DB" "$BACKUP_DIR/$BACKUP_NAME"
    echo -e "${GREEN}✓ Local backup created: $BACKUP_DIR/$BACKUP_NAME${NC}"

    # Pull remote to local
    echo -e "${BLUE}⬇ Pulling remote database to local...${NC}"
    rsync -avz "$SERVER:$REMOTE_DB" "$LOCAL_DB"
    echo -e "${GREEN}✓ Database pulled successfully${NC}"
fi

echo ""
echo "========================================="
echo -e "${GREEN}✓ Database sync complete!${NC}"
echo "========================================="
echo ""
echo "Backups are stored in:"
if [ "$SYNC_DIRECTION" = "push" ]; then
    echo "  Remote: /home/admin/portfolio/.db-backups/"
else
    echo "  Local: $BACKUP_DIR/"
fi
echo ""
