#!/usr/bin/env bash
# Complete deployment: Application + Caddy Configuration

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "========================================="
echo "Full Deployment to michaelhegner.com"
echo "========================================="

# Step 1: Deploy the application
echo ""
echo "STEP 1: Deploying application..."
"$SCRIPT_DIR/deploy.sh"

# Step 2: Update Caddy configuration
echo ""
echo "STEP 2: Updating Caddy configuration..."
"$SCRIPT_DIR/update-caddy.sh"

echo ""
echo "========================================="
echo "✓ FULL DEPLOYMENT COMPLETE!"
echo "========================================="
echo ""
echo "Your portfolio is now live at:"
echo "  https://michaelhegner.com"
echo ""
echo "Changes may take a moment to propagate."
echo "If you see a Caddy error, the SSL certificate"
echo "is being issued and should be ready in ~30 seconds."
echo ""
