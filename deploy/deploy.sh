#!/usr/bin/env bash
# Deployment script for Portfolio V2

set -e

SERVER="admin@50.116.26.167"
REMOTE_DIR="/home/admin/portfolio"
LOCAL_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "========================================="
echo "Deploying Portfolio V2 to Linode"
echo "========================================="

# Step 0: Check/Set Admin Password
echo ""
echo "Step 0: Admin Authentication Setup..."

# Try to load from .env file first
if [ -f "$LOCAL_DIR/.env" ] && [ -z "$ADMIN_PASSWORD" ]; then
    echo "Loading credentials from .env file..."
    export $(grep -v '^#' "$LOCAL_DIR/.env" | grep ADMIN_PASSWORD | xargs)
fi

if [ -z "$ADMIN_PASSWORD" ]; then
    echo "ADMIN_PASSWORD not found in .env file or environment."
    echo "Please enter a secure password for admin access:"
    read -s -p "Admin Password: " ADMIN_PASSWORD
    echo ""
    read -s -p "Confirm Password: " ADMIN_PASSWORD_CONFIRM
    echo ""

    if [ "$ADMIN_PASSWORD" != "$ADMIN_PASSWORD_CONFIRM" ]; then
        echo "❌ Passwords do not match. Deployment aborted."
        exit 1
    fi

    if [ ${#ADMIN_PASSWORD} -lt 10 ]; then
        echo "❌ Password must be at least 10 characters. Deployment aborted."
        exit 1
    fi
fi
echo "✓ Admin password configured"

# Step 0.5: Sync database (ensure we're deploying the latest content)
echo ""
echo "Step 0.5: Syncing database..."
if [ -f "$LOCAL_DIR/scripts/sync-db.sh" ]; then
    "$LOCAL_DIR/scripts/sync-db.sh"
else
    echo "⚠ Database sync script not found, skipping..."
fi

# Step 1: Build the application
echo ""
echo "Step 1: Building production binary for Linux..."
cd "$LOCAL_DIR"
templ generate
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o portfolio-v2 main.go
echo "✓ Build complete"

# Step 2: Create deployment package
echo ""
echo "Step 2: Creating deployment package..."
mkdir -p deploy/package
cp portfolio-v2 deploy/package/
cp portfolio.db deploy/package/
cp -r static deploy/package/
echo "✓ Package created"

# Step 3: Transfer to server
echo ""
echo "Step 3: Transferring files to server..."
echo "Creating remote directory..."
ssh "$SERVER" "mkdir -p $REMOTE_DIR"

echo "Syncing files with rsync..."
rsync -avz --delete \
    --exclude='.DS_Store' \
    deploy/package/ "$SERVER:$REMOTE_DIR/"

# Update service file with admin password
echo "Configuring admin credentials..."
sed "s/CHANGE_THIS_PASSWORD/$ADMIN_PASSWORD/g" deploy/portfolio.service > /tmp/portfolio.service.tmp
rsync -avz /tmp/portfolio.service.tmp "$SERVER:/tmp/portfolio.service"
rm /tmp/portfolio.service.tmp
echo "✓ Files transferred"

# Step 4: Set up systemd service
echo ""
echo "Step 4: Setting up systemd service..."

# Set permissions first (no sudo needed)
ssh "$SERVER" "chmod +x /home/admin/portfolio/portfolio-v2"

# Move service file and configure systemd (with passwordless sudo)
ssh "$SERVER" "sudo mv /tmp/portfolio.service /etc/systemd/system/portfolio.service && \
    sudo systemctl daemon-reload && \
    sudo systemctl enable portfolio && \
    sudo systemctl restart portfolio && \
    sleep 2 && \
    sudo systemctl status portfolio --no-pager"

echo "✓ Service configured"

# Cleanup
echo ""
echo "Cleaning up local deployment package..."
rm -rf deploy/package
echo "✓ Cleanup complete"

echo ""
echo "========================================="
echo "✓ Deployment Complete!"
echo "========================================="
echo ""
echo "Your portfolio is now running on:"
echo "http://50.116.26.167:8080"
echo ""
echo "Useful commands:"
echo "  Check status:  ssh $SERVER 'sudo systemctl status portfolio'"
echo "  View logs:     ssh $SERVER 'sudo journalctl -u portfolio -f'"
echo "  Restart:       ssh $SERVER 'sudo systemctl restart portfolio'"
echo "  Stop:          ssh $SERVER 'sudo systemctl stop portfolio'"
echo ""
