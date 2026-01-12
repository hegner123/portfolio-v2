# Admin Authentication Guide

## Overview

Your portfolio now uses **HTTP Basic Authentication** to protect admin routes:
- `/admin/blog/new` - Create new blog posts
- `/admin/project/new` - Create new projects

## How It Works

When you visit an admin page, your browser will prompt you to enter:
- **Username**: `admin`
- **Password**: The password you set during deployment

The browser will remember your credentials for the session.

## Setting Up Admin Access

### Option 1: During Deployment (Recommended)

When you run the deployment script:

```bash
./deploy/deploy.sh
```

The script will prompt you to set an admin password:

```
Step 0: Admin Authentication Setup...
ADMIN_PASSWORD environment variable not set.
Please enter a secure password for admin access:
Admin Password: ************
Confirm Password: ************
✓ Admin password configured
```

**Password Requirements**:
- Minimum 12 characters
- Should be strong and unique
- Will be stored securely on the server

### Option 2: Set Environment Variable

If you want to skip the prompt, set the password before deploying:

```bash
export ADMIN_PASSWORD="your-secure-password-here"
./deploy/deploy.sh
```

## Testing Locally

To test authentication on your local machine:

```bash
# Set credentials
export ADMIN_USERNAME="admin"
export ADMIN_PASSWORD="testpassword123"

# Run the application
go run main.go
```

Visit: http://localhost:8080/admin/blog/new

Your browser will prompt for credentials. Enter:
- Username: `admin`
- Password: `testpassword123`

## Accessing Admin Pages

### First Time Access

1. Visit: https://michaelhegner.com/admin/blog/new
2. Browser will show authentication dialog
3. Enter credentials:
   - Username: `admin`
   - Password: [your password]
4. Click "Sign In" or "Log In"

### Subsequent Access

Your browser will remember the credentials for the current session. You won't need to re-enter them unless you:
- Close the browser
- Clear cookies/cache
- Use incognito/private mode

## Logging Out

To log out of admin access:
1. Close all browser windows
2. Or clear your browser's authentication cache
3. Or use Incognito/Private mode for temporary access

### Chrome/Edge:
- Settings → Privacy → Clear browsing data → Passwords

### Firefox:
- Settings → Privacy → History → Clear Recent History → Active Logins

### Safari:
- Safari → Preferences → Privacy → Manage Website Data

## Changing Admin Password

To change your admin password:

1. Update the password on the server:

```bash
ssh admin@50.116.26.167
sudo nano /etc/systemd/system/portfolio.service
```

2. Change the `ADMIN_PASSWORD` line:

```ini
Environment="ADMIN_PASSWORD=your-new-password"
```

3. Reload and restart the service:

```bash
sudo systemctl daemon-reload
sudo systemctl restart portfolio
```

Or simply redeploy with a new password:

```bash
export ADMIN_PASSWORD="new-password"
./deploy/deploy.sh
```

## Security Features

✅ **Password Protection**: Admin routes require authentication
✅ **Logging**: Failed login attempts are logged
✅ **HTTPS**: Credentials encrypted in transit (via Caddy)
✅ **Server-Side Validation**: Password never exposed to client

## Troubleshooting

### "Unauthorized" Every Time

**Cause**: Wrong username or password

**Solution**:
1. Check that you're using username `admin`
2. Verify password is correct
3. Check server logs: `ssh admin@50.116.26.167 'sudo journalctl -u portfolio -n 50'`

### Browser Not Prompting for Password

**Cause**: Browser has cached old credentials

**Solution**:
1. Clear authentication cache (see "Logging Out" above)
2. Try Incognito/Private mode
3. Try a different browser

### "Admin access not configured"

**Cause**: Environment variables not set on server

**Solution**:
1. Check the service file has credentials:
   ```bash
   ssh admin@50.116.26.167 'sudo cat /etc/systemd/system/portfolio.service | grep ADMIN'
   ```
2. Redeploy with password set

### Password in Browser But Still Prompts

**Cause**: Password may have changed on server

**Solution**:
1. Clear saved password in browser
2. Enter the current password

## Best Practices

### ✅ Do:
- Use a strong, unique password (12+ characters)
- Use a password manager to store it
- Change password periodically
- Use HTTPS only (already configured)
- Monitor failed login attempts in logs

### ❌ Don't:
- Share your admin password
- Use the same password as other services
- Store password in plaintext notes
- Access admin pages on public WiFi without VPN
- Commit the password to git (already prevented)

## Next Steps (Optional Upgrades)

Current implementation: **Phase 1 - HTTP Basic Auth**

To upgrade security:

### Phase 2: Session-Based Auth
- Proper login/logout UI
- Session management
- CSRF protection
- Better UX

See `/ai/ADMIN-SECURITY-IMPLEMENTATION.md` for upgrade guide.

## Getting Help

If you have issues:
1. Check the troubleshooting section above
2. Review server logs: `ssh admin@50.116.26.167 'sudo journalctl -u portfolio -f'`
3. Test locally with environment variables set
4. Verify Caddy is running: `ssh admin@50.116.26.167 'sudo systemctl status caddy'`
