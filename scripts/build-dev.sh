#!/usr/bin/env bash

# Development Build Script for Portfolio V2
# Quick build for local development and testing

set -e  # Exit on error

echo "🛠️  Starting development build..."

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Step 1: Generate Templ templates
echo -e "${BLUE}📝 Generating Templ templates...${NC}"
templ generate
echo -e "${GREEN}✓ Templates generated${NC}"

# Step 2: Build Go binary
echo -e "${BLUE}🔨 Building Go binary...${NC}"
go build -o portfolio-v2 main.go
echo -e "${GREEN}✓ Binary built${NC}"

echo ""
echo -e "${GREEN}✓ Development build complete!${NC}"
echo -e "${BLUE}Run the server: ./portfolio-v2${NC}"
echo -e "${BLUE}Or use: go run main.go${NC}"
echo ""
