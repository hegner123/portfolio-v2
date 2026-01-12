#!/usr/bin/env bash
# Update Caddy configuration for Portfolio V2

set -e

SERVER="admin@50.116.26.167"

echo "========================================="
echo "Updating Caddy Configuration"
echo "========================================="

# Step 1: Backup existing Caddyfile
echo ""
echo "Step 1: Backing up existing Caddyfile..."
ssh "$SERVER" "sudo cp /etc/caddy/Caddyfile /etc/caddy/Caddyfile.backup.$(date +%Y%m%d_%H%M%S) 2>/dev/null || true"
echo "✓ Backup complete"

# Step 2: Upload new Caddyfile
echo ""
echo "Step 2: Uploading new Caddyfile..."
rsync -avz deploy/Caddyfile "$SERVER:/tmp/Caddyfile"
echo "✓ Upload complete"

# Step 3: Move to Caddy directory and reload
echo ""
echo "Step 3: Installing and reloading Caddy..."
ssh "$SERVER" "sudo mv /tmp/Caddyfile /etc/caddy/Caddyfile && \
    sudo chown root:root /etc/caddy/Caddyfile && \
    sudo chmod 644 /etc/caddy/Caddyfile && \
    echo 'Validating Caddy configuration...' && \
    sudo caddy validate --config /etc/caddy/Caddyfile && \
    echo 'Reloading Caddy...' && \
    sudo systemctl reload caddy && \
    sleep 2 && \
    sudo systemctl status caddy --no-pager"
echo "✓ Caddy reloaded"

echo ""
echo "========================================="
echo "✓ Caddy Configuration Updated!"
echo "========================================="
echo ""
echo "Your portfolio will be accessible at:"
echo "  https://michaelhegner.com"
echo "  https://www.michaelhegner.com (redirects to non-www)"
echo ""
echo "HTTPS is automatically enabled via Let's Encrypt."
echo ""
echo "Useful commands:"
echo "  Check Caddy status:  ssh $SERVER 'sudo systemctl status caddy'"
echo "  View Caddy logs:     ssh $SERVER 'sudo journalctl -u caddy -f'"
echo "  Reload config:       ssh $SERVER 'sudo systemctl reload caddy'"
echo ""
