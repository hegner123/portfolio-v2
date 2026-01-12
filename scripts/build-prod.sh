#!/usr/bin/env bash

# Production Build Script for Portfolio V2
# This script optimizes the application for production deployment

set -e  # Exit on error

echo "🚀 Starting production build..."

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Step 1: Generate Templ templates
echo -e "${BLUE}📝 Generating Templ templates...${NC}"
templ generate
echo -e "${GREEN}✓ Templates generated${NC}"

# Step 2: Tidy Go dependencies
echo -e "${BLUE}🧹 Tidying Go modules...${NC}"
go mod tidy
echo -e "${GREEN}✓ Dependencies tidied${NC}"

# Step 3: Build Go binary with optimizations
echo -e "${BLUE}🔨 Building optimized Go binary for Linux...${NC}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o portfolio-v2 main.go
echo -e "${GREEN}✓ Binary built (portfolio-v2)${NC}"

# Step 4: Check if CSS minification tools are available
echo -e "${BLUE}📦 Checking for CSS optimization tools...${NC}"

# Check if we should minify CSS
MINIFY_CSS=false

# Try to find csso-cli (npm package)
if command -v csso &> /dev/null; then
    MINIFY_CSS=true
    MINIFIER="csso"
    echo -e "${GREEN}✓ Found csso for CSS minification${NC}"
# Try to find clean-css-cli (npm package)
elif command -v cleancss &> /dev/null; then
    MINIFY_CSS=true
    MINIFIER="cleancss"
    echo -e "${GREEN}✓ Found clean-css for CSS minification${NC}"
else
    echo -e "${BLUE}ℹ CSS minification skipped (install csso-cli or clean-css-cli for minification)${NC}"
    echo -e "${BLUE}  npm install -g csso-cli${NC}"
fi

# Step 5: Minify CSS if tools available
if [ "$MINIFY_CSS" = true ]; then
    echo -e "${BLUE}🎨 Minifying CSS files...${NC}"

    # Create minified directory if it doesn't exist
    mkdir -p static/css/min

    # Minify each CSS file
    for css_file in static/css/*.css; do
        # Skip files in the min directory
        if [[ "$css_file" != *"/min/"* ]]; then
            filename=$(basename "$css_file")

            if [ "$MINIFIER" = "csso" ]; then
                csso "$css_file" -o "static/css/min/$filename"
            elif [ "$MINIFIER" = "cleancss" ]; then
                cleancss -o "static/css/min/$filename" "$css_file"
            fi

            # Calculate size reduction
            original_size=$(wc -c < "$css_file" | tr -d ' ')
            minified_size=$(wc -c < "static/css/min/$filename" | tr -d ' ')
            reduction=$((100 - (minified_size * 100 / original_size)))

            echo -e "${GREEN}  ✓ $filename (${reduction}% smaller)${NC}"
        fi
    done

    echo -e "${GREEN}✓ CSS minified to static/css/min/${NC}"
    echo -e "${BLUE}ℹ To use minified CSS in production, update layout.templ to reference /static/css/min/ files${NC}"
fi

# Step 6: Display build summary
echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}✓ Production build complete!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "${BLUE}Build artifacts:${NC}"
echo -e "  • Binary: ./portfolio-v2"
echo -e "  • Size: $(du -h portfolio-v2 | cut -f1)"
if [ "$MINIFY_CSS" = true ]; then
    echo -e "  • Minified CSS: ./static/css/min/"
fi
echo ""
echo -e "${BLUE}Next steps:${NC}"
echo -e "  1. Test the build: ./portfolio-v2"
echo -e "  2. Deploy to your server"
echo -e "  3. Don't forget to copy the static/ directory"
echo ""

# Step 7: Optional - Create deployment package
read -p "Create deployment tarball? (y/N) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${BLUE}📦 Creating deployment package...${NC}"
    tar -czf portfolio-v2-deploy.tar.gz portfolio-v2 static/ portfolio.db
    echo -e "${GREEN}✓ Created portfolio-v2-deploy.tar.gz${NC}"
    echo -e "  Size: $(du -h portfolio-v2-deploy.tar.gz | cut -f1)"
fi

echo -e "${GREEN}🎉 Done!${NC}"
