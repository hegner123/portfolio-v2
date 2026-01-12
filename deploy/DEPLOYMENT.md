# Deployment Guide

## Quick Deployment (Automated)

### Full Deployment (Application + Caddy)

Run this single command to deploy everything:

```bash
./deploy/full-deploy.sh
```

**Note**: The script will prompt you to set an admin password for accessing admin pages (`/admin/blog/new`, `/admin/project/new`).

This will deploy the application AND update Caddy to serve it at https://michaelhegner.com

### Admin Authentication

After deployment:
- Visit: https://michaelhegner.com/admin/blog/new or https://michaelhegner.com/admin/project/new
- Your browser will prompt for credentials
- Username: `admin`
- Password: The password you set during deployment

See `/deploy/ADMIN-AUTH-GUIDE.md` for complete authentication guide.

### Deploy Application Only

```bash
./deploy/deploy.sh
```

This script will:
1. Build the production binary
2. Create a deployment package
3. Transfer files to the server
4. Set up and start the systemd service
5. Show the service status

### Update Caddy Configuration Only

```bash
./deploy/update-caddy.sh
```

This will update the Caddy configuration to reverse proxy to your application with automatic HTTPS

## Manual Deployment (Step-by-Step)

If you prefer to deploy manually or troubleshoot:

### 1. Build the Application

```bash
templ generate
go build -ldflags="-s -w" -o portfolio-v2 main.go
```

### 2. Transfer Files to Server

```bash
# Create directory on server
ssh admin@50.116.26.167 "mkdir -p /home/admin/portfolio"

# Upload application files
scp portfolio-v2 admin@50.116.26.167:/home/admin/portfolio/
scp portfolio.db admin@50.116.26.167:/home/admin/portfolio/
scp -r static admin@50.116.26.167:/home/admin/portfolio/

# Upload service file
scp deploy/portfolio.service admin@50.116.26.167:/tmp/
```

### 3. Set Up the Service

SSH into the server:

```bash
ssh admin@50.116.26.167
```

Then run these commands on the server:

```bash
# Make binary executable
chmod +x /home/admin/portfolio/portfolio-v2

# Install systemd service
sudo mv /tmp/portfolio.service /etc/systemd/system/portfolio.service

# Reload systemd
sudo systemctl daemon-reload

# Enable service to start on boot
sudo systemctl enable portfolio

# Start the service
sudo systemctl start portfolio

# Check status
sudo systemctl status portfolio
```

### 4. Verify Deployment

Visit: http://50.116.26.167:8080

## Useful Commands

### Check Service Status
```bash
ssh admin@50.116.26.167 'sudo systemctl status portfolio'
```

### View Live Logs
```bash
ssh admin@50.116.26.167 'sudo journalctl -u portfolio -f'
```

### Restart Service
```bash
ssh admin@50.116.26.167 'sudo systemctl restart portfolio'
```

### Stop Service
```bash
ssh admin@50.116.26.167 'sudo systemctl stop portfolio'
```

### View Last 100 Log Lines
```bash
ssh admin@50.116.26.167 'sudo journalctl -u portfolio -n 100 --no-pager'
```

## Setting Up Domain & HTTPS (Optional)

If you want to use a custom domain with HTTPS:

### 1. Install Nginx

```bash
ssh admin@50.116.26.167
sudo apt update
sudo apt install nginx
```

### 2. Configure Nginx Reverse Proxy

Create `/etc/nginx/sites-available/portfolio`:

```nginx
server {
    listen 80;
    server_name yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Enable the site:

```bash
sudo ln -s /etc/nginx/sites-available/portfolio /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

### 3. Install SSL with Certbot

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com
```

## Troubleshooting

### Service won't start

Check logs:
```bash
ssh admin@50.116.26.167 'sudo journalctl -u portfolio -n 50'
```

### Permission denied

Ensure binary is executable:
```bash
ssh admin@50.116.26.167 'chmod +x /home/admin/portfolio/portfolio-v2'
```

### Port already in use

Check what's using port 8080:
```bash
ssh admin@50.116.26.167 'sudo lsof -i :8080'
```

### Database errors

Ensure database file exists and has correct permissions:
```bash
ssh admin@50.116.26.167 'ls -la /home/admin/portfolio/portfolio.db'
ssh admin@50.116.26.167 'chmod 644 /home/admin/portfolio/portfolio.db'
```

## Updating the Application

To deploy updates, simply run the deployment script again:

```bash
./deploy/deploy.sh
```

This will rebuild, transfer, and restart the service automatically.
