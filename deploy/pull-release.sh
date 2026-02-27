#!/usr/bin/env bash
# pull-release.sh - Download latest GitHub release and deploy
set -euo pipefail

REPO="hegner123/portfolio-v2"
DEPLOY_DIR="/home/admin/portfolio"
SERVICE="portfolio"
TARBALL="/tmp/portfolio-v2.tar.gz"

echo "Fetching latest release from ${REPO}..."
DOWNLOAD_URL=$(curl -sf "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep -o '"browser_download_url": *"[^"]*portfolio-v2\.tar\.gz"' \
    | head -1 \
    | cut -d'"' -f4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "ERROR: Could not find release tarball URL" >&2
    exit 1
fi

TAG=$(curl -sf "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep -o '"tag_name": *"[^"]*"' \
    | head -1 \
    | cut -d'"' -f4)
echo "Latest release: ${TAG}"

echo "Downloading ${DOWNLOAD_URL}..."
curl -sfL -o "$TARBALL" "$DOWNLOAD_URL"

echo "Stopping ${SERVICE}..."
sudo systemctl stop "$SERVICE"

echo "Backing up current binary..."
if [ -f "${DEPLOY_DIR}/portfolio-v2" ]; then
    cp "${DEPLOY_DIR}/portfolio-v2" "${DEPLOY_DIR}/portfolio-v2.bak"
fi

echo "Extracting release..."
tar xzf "$TARBALL" -C "$DEPLOY_DIR"

echo "Setting permissions..."
chmod +x "${DEPLOY_DIR}/portfolio-v2"

echo "Starting ${SERVICE}..."
sudo systemctl start "$SERVICE"

echo "Cleaning up..."
rm -f "$TARBALL"

echo "Checking service status..."
sleep 2
sudo systemctl status "$SERVICE" --no-pager

echo ""
echo "Deploy complete: ${TAG}"
