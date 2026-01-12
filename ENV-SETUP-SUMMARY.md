# Environment Variable Setup Summary

## Overview

Added .env file support and a user-friendly setup page that displays when admin credentials are not configured.

## Changes Made

### 1. Added .env File Support

**Modified Files:**
- `main.go` - Added godotenv to load .env on startup
- `go.mod` - Added github.com/joho/godotenv dependency

**New Files:**
- `.env.example` - Template for environment variables

### 2. Created Admin Setup Page

When credentials are not configured, users now see a helpful setup page instead of a plain error.

**New Files:**
- `templates/admin-setup.templ` - Setup instructions page
- `static/css/admin-setup.css` - Styling for setup page

**Modified Files:**
- `middleware/auth.go` - Renders setup page when credentials missing
- `templates/layout.templ` - Added admin-setup.css

## How It Works

### Startup Process

1. Application starts
2. `godotenv.Load()` attempts to load `.env` file
3. If `.env` doesn't exist, falls back to system environment variables
4. Server starts normally

### Admin Access Flow

1. User visits `/admin` (or any admin route)
2. Middleware checks for `ADMIN_USERNAME` and `ADMIN_PASSWORD`
3. **If NOT configured:**
   - Shows beautiful setup page with step-by-step instructions
   - Returns 503 Service Unavailable status
4. **If configured:**
   - Shows HTTP Basic Auth login prompt
   - Validates credentials
   - Grants or denies access

## Setup Instructions for Users

### Quick Start

1. **Copy the example file:**
   ```bash
   cp .env.example .env
   ```

2. **Edit `.env` with your credentials:**
   ```env
   ADMIN_USERNAME=admin
   ADMIN_PASSWORD=your-secure-password-here
   PORT=8080
   ```

3. **Start the server:**
   ```bash
   go run main.go
   ```

4. **Visit `/admin`:**
   - Browser will prompt for credentials
   - Enter the username and password from your `.env` file
   - Access granted!

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `ADMIN_USERNAME` | Yes | - | Username for admin access |
| `ADMIN_PASSWORD` | Yes | - | Password for admin access |
| `PORT` | No | 8080 | Server port |

## Security Features

### .env File Protection

**IMPORTANT:** Add `.env` to your `.gitignore`:

```gitignore
# Environment variables
.env
```

This prevents accidentally committing sensitive credentials to version control.

### Password Requirements

For production, use a strong password:
- Minimum 12 characters
- Mix of uppercase, lowercase, numbers, and symbols
- Avoid common words or patterns

Example strong password:
```
Tr0ub4dor&3_SecureP@ss!2026
```

## Setup Page Features

The admin setup page includes:

✅ **Clear visual design** - Lock icon and gradient title
✅ **Step-by-step instructions** - Numbered steps with code examples
✅ **Copy-paste ready** - Pre-formatted .env template
✅ **Security reminder** - Warning about .gitignore
✅ **Responsive design** - Works on all devices

### Visual Preview

```
┌─────────────────────────────────────────────┐
│                    🔒                        │
│        Admin Access Not Configured          │
│                                             │
│  Admin credentials are not set up yet...   │
│                                             │
│  ┌───────────────────────────────────────┐ │
│  │ 1  Create a .env file                 │ │
│  │    touch .env                         │ │
│  └───────────────────────────────────────┘ │
│                                             │
│  ┌───────────────────────────────────────┐ │
│  │ 2  Add your credentials               │ │
│  │    ADMIN_USERNAME=admin               │ │
│  │    ADMIN_PASSWORD=secure-pass         │ │
│  └───────────────────────────────────────┘ │
│                                             │
│  ┌───────────────────────────────────────┐ │
│  │ 3  Restart the server                 │ │
│  │    go run main.go                     │ │
│  └───────────────────────────────────────┘ │
│                                             │
│  ┌───────────────────────────────────────┐ │
│  │ 4  Access the admin dashboard         │ │
│  │    Visit /admin                       │ │
│  └───────────────────────────────────────┘ │
│                                             │
│  ⚠️ Add .env to .gitignore                 │
│                                             │
│           [← Back to Home]                  │
└─────────────────────────────────────────────┘
```

## Deployment

### Local Development

1. Create `.env` file
2. Run: `go run main.go`
3. Visit: `http://localhost:8080/admin`

### Production (Linode Server)

The deployment script (`deploy/deploy.sh`) already handles environment variables:

1. During deployment, you're prompted for admin password
2. Script sets environment variables on the server
3. Systemd service loads them automatically

**Deployment command:**
```bash
./deploy/full-deploy.sh
```

## Troubleshooting

### Issue: "Admin access not configured" page appears

**Solution:**
1. Check if `.env` file exists in project root
2. Verify `ADMIN_USERNAME` and `ADMIN_PASSWORD` are set in `.env`
3. Restart the server
4. Try accessing `/admin` again

### Issue: .env file not loading

**Check:**
- File is named exactly `.env` (not `.env.txt`)
- File is in the same directory as `main.go`
- File permissions are readable
- No syntax errors in .env file

**Debug:**
```bash
# Verify .env exists
ls -la .env

# Check contents
cat .env

# Verify environment variables are loaded
go run main.go
# Look for: "No .env file found" vs silent loading
```

### Issue: Still being denied after setting credentials

**Check:**
- Credentials in `.env` match what you're entering
- No extra spaces or quotes around values
- Server was restarted after creating `.env`

## Benefits

### Developer Experience

- ✅ **No command-line exports needed** - Just edit .env
- ✅ **Version controlled template** - .env.example in repo
- ✅ **Clear error messages** - Beautiful setup page with instructions
- ✅ **Fast onboarding** - New developers can set up in 2 minutes

### Security

- ✅ **Credentials in files, not code** - Environment variable best practice
- ✅ **Git-ignored by default** - Won't accidentally commit secrets
- ✅ **Production-ready** - Same pattern works for deployment

### User-Friendly

- ✅ **Visual guidance** - Step-by-step instructions with code blocks
- ✅ **No confusion** - Clear next steps when not configured
- ✅ **Responsive design** - Setup page works on any device

---

**Status**: ✅ Complete! .env support and setup page fully functional.

**Next**: Create your `.env` file and start managing content!
