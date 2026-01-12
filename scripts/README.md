# Build Scripts

This directory contains build scripts for the Portfolio V2 project.

## Available Scripts

### `build-dev.sh` - Development Build

Quick build script for local development and testing.

**Usage:**
```bash
./scripts/build-dev.sh
```

**What it does:**
1. Generates Templ templates
2. Builds Go binary without optimizations

**Output:**
- `./portfolio-v2` - Development binary

---

### `build-prod.sh` - Production Build

Comprehensive build script with optimizations for production deployment.

**Usage:**
```bash
./scripts/build-prod.sh
```

**What it does:**
1. Generates Templ templates
2. Tidies Go modules
3. Builds optimized Go binary (with `-ldflags="-s -w"`)
4. Minifies CSS files (if tools available)
5. Optionally creates deployment tarball

**Output:**
- `./portfolio-v2` - Production-optimized binary
- `./static/css/min/` - Minified CSS files (if minifier available)
- `./portfolio-v2-deploy.tar.gz` - Deployment package (optional)

**CSS Minification:**

The script automatically detects and uses available CSS minification tools:
- `csso-cli` (recommended)
- `clean-css-cli`

To install a minifier:
```bash
# Option 1: csso (recommended)
npm install -g csso-cli

# Option 2: clean-css
npm install -g clean-css-cli
```

**Note:** CSS minification is optional. The build will succeed without it, but CSS files won't be minified.

---

## Build Flags Explained

### Production Build Flags

`go build -ldflags="-s -w"`

- `-s`: Strip symbol table (reduces binary size)
- `-w`: Strip DWARF debugging information (reduces binary size)

These flags significantly reduce the binary size, making it ideal for production deployment. Debug symbols are removed, so this binary is not suitable for debugging.

### Development Build

Uses default `go build` without optimization flags, keeping debug symbols for easier development and debugging.

---

## Performance Optimizations

The production build includes several performance optimizations:

1. **Binary Size Reduction**: Strip flags reduce binary size by 20-30%
2. **CSS Minification**: Removes whitespace, comments, and optimizes CSS (if tools available)
3. **Module Tidying**: Ensures only necessary dependencies are included

---

## Deployment

### Manual Deployment

After running `build-prod.sh`:

```bash
# 1. Build for production
./scripts/build-prod.sh

# 2. Transfer to server
scp portfolio-v2 user@server:/path/to/app/
scp -r static user@server:/path/to/app/
scp portfolio.db user@server:/path/to/app/

# 3. Restart service on server
ssh user@server "systemctl restart portfolio"
```

### Using Deployment Tarball

If you created a deployment package:

```bash
# 1. Transfer tarball
scp portfolio-v2-deploy.tar.gz user@server:/tmp/

# 2. Extract on server
ssh user@server
cd /path/to/app
tar -xzf /tmp/portfolio-v2-deploy.tar.gz

# 3. Restart service
systemctl restart portfolio
```

---

## Troubleshooting

### Templ CLI Not Found

```bash
# Install Templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Ensure $GOPATH/bin is in your PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### CSS Minification Skipped

This is normal if you don't have a CSS minifier installed. The build will complete successfully, but CSS won't be minified.

To enable CSS minification:
```bash
npm install -g csso-cli
```

### Permission Denied

Make sure the scripts are executable:
```bash
chmod +x scripts/build-dev.sh
chmod +x scripts/build-prod.sh
```

---

## Advanced Usage

### Custom Build Flags

You can modify the build scripts or run custom builds:

```bash
# Build with race detector (development/testing)
go build -race -o portfolio-v2 main.go

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o portfolio-v2-linux main.go

# Build with version info
VERSION=$(git describe --tags --always)
go build -ldflags="-s -w -X main.Version=$VERSION" -o portfolio-v2 main.go
```

### CI/CD Integration

These scripts are designed to be easily integrated into CI/CD pipelines:

```yaml
# Example GitHub Actions workflow
- name: Build for production
  run: ./scripts/build-prod.sh

- name: Upload artifact
  uses: actions/upload-artifact@v2
  with:
    name: portfolio-v2
    path: portfolio-v2
```

---

## See Also

- [CLAUDE.md](/CLAUDE.md) - Development workflow and commands
- [ai/IMPLEMENTATION.md](/ai/IMPLEMENTATION.md) - Implementation tracker
- [ai/conventions.md](/ai/conventions.md) - Coding conventions
