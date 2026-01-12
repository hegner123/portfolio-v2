# Admin Routes Security Implementation Plan

## Current Status
- ❌ Admin routes are publicly accessible
- ❌ No authentication required
- ❌ No authorization checks
- Routes: `/admin/blog/new`, `/admin/project/new`

---

## Implementation Phases

### Phase 1: Basic HTTP Basic Auth (Quick & Simple)
**Status**: Not Started
**Estimated Time**: 30 minutes
**Security Level**: Basic (⭐⭐☆☆☆)

**What to Implement**:
- HTTP Basic Authentication for all `/admin/*` routes
- Middleware to check credentials
- Single admin username/password stored in environment variable

**Tasks**:
- [ ] Create auth middleware function
- [ ] Add environment variables for ADMIN_USERNAME and ADMIN_PASSWORD
- [ ] Wrap admin routes with auth middleware
- [ ] Test authentication flow
- [ ] Update deployment with environment variables

**Pros**:
- Very simple to implement
- No database changes needed
- Works immediately
- Built into HTTP standard

**Cons**:
- Credentials sent with every request
- No session management
- Password visible in browser password manager
- Single user only
- No logout functionality

**Implementation Files**:
- `middleware/auth.go` - Basic auth middleware
- `.env` or environment config
- Update `main.go` to use middleware

---

### Phase 2: Session-Based Authentication (Recommended)
**Status**: Not Started
**Estimated Time**: 2-3 hours
**Security Level**: Good (⭐⭐⭐⭐☆)

**What to Implement**:
- Login page with form
- Session management with secure cookies
- Password hashing (bcrypt)
- CSRF protection
- Session expiration
- Logout functionality

**Tasks**:
- [ ] Create users table in database
- [ ] Hash and store admin password
- [ ] Create login page template
- [ ] Implement session store (in-memory or database)
- [ ] Create login/logout handlers
- [ ] Create session middleware
- [ ] Add CSRF tokens to forms
- [ ] Implement "remember me" functionality (optional)
- [ ] Add session timeout (e.g., 24 hours)
- [ ] Update all admin pages with logout button

**Database Schema**:
```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    last_login DATETIME
);

CREATE TABLE IF NOT EXISTS sessions (
    id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

**Pros**:
- Secure password storage
- Better UX (login once, stay logged in)
- Support for logout
- Session expiration
- Can support multiple users
- CSRF protection

**Cons**:
- More complex to implement
- Requires database changes
- Session storage management

**Implementation Files**:
- `database/users.go` - User CRUD operations
- `database/sessions.go` - Session management
- `handlers/auth.go` - Login/logout handlers
- `middleware/session.go` - Session validation
- `templates/login.templ` - Login page
- `models/user.go` - User model
- `models/session.go` - Session model

---

### Phase 3: Token-Based Auth with JWT (API-First)
**Status**: Not Started
**Estimated Time**: 3-4 hours
**Security Level**: Good (⭐⭐⭐⭐☆)

**What to Implement**:
- JWT token generation and validation
- Refresh token mechanism
- Token storage (httpOnly cookies or localStorage)
- Token expiration and refresh

**Tasks**:
- [ ] Install JWT library (github.com/golang-jwt/jwt)
- [ ] Create token generation functions
- [ ] Create token validation middleware
- [ ] Implement refresh token logic
- [ ] Store tokens in httpOnly cookies
- [ ] Add token expiration (15-60 minutes)
- [ ] Add refresh token expiration (7-30 days)
- [ ] Create login endpoint returning JWT
- [ ] Update frontend to handle tokens

**Pros**:
- Stateless (no server-side session storage)
- Scales well
- Works great for APIs
- Can be used across domains

**Cons**:
- Cannot revoke tokens before expiry (without blacklist)
- More complex implementation
- Requires careful secret management
- Token refresh adds complexity

**Implementation Files**:
- `auth/jwt.go` - JWT generation/validation
- `middleware/jwt.go` - JWT middleware
- `handlers/auth.go` - Token endpoints
- Environment variables for JWT_SECRET

---

### Phase 4: OAuth2/Social Login (Advanced)
**Status**: Not Started
**Estimated Time**: 4-6 hours
**Security Level**: Excellent (⭐⭐⭐⭐⭐)

**What to Implement**:
- GitHub OAuth integration
- Google OAuth integration
- User whitelist (only allow specific GitHub/Google accounts)

**Tasks**:
- [ ] Register OAuth apps (GitHub/Google)
- [ ] Install OAuth library
- [ ] Create OAuth callback handlers
- [ ] Store OAuth tokens
- [ ] Implement user whitelist
- [ ] Create user profile page
- [ ] Link OAuth accounts to local users

**Pros**:
- No password management
- Users already authenticated by trusted provider
- Can restrict to specific accounts
- Professional authentication

**Cons**:
- Depends on third-party services
- More complex setup
- Requires OAuth app registration
- Potential privacy concerns

**Implementation Files**:
- `auth/oauth.go` - OAuth handlers
- `handlers/oauth-callback.go` - OAuth callbacks
- Configuration for OAuth credentials
- User whitelist in database or config

---

## Recommended Approach

### For Personal Portfolio (Single User)
**Choose Phase 1 or Phase 2**

#### Start with Phase 1 (HTTP Basic Auth) if:
- You need security NOW
- You're the only admin
- You want minimal complexity
- You don't mind re-entering password

#### Move to Phase 2 (Session-Based) if:
- You want better UX
- You plan to add more admins later
- You want proper logout functionality
- You want sessions to persist

### For Multi-User or Production Application
**Go directly to Phase 2 or Phase 3**

---

## Security Best Practices (All Phases)

### Required Security Measures:
- [ ] **HTTPS Only**: All admin routes must use HTTPS in production
- [ ] **Rate Limiting**: Limit login attempts (5 attempts per 15 minutes)
- [ ] **Secure Cookies**: Use httpOnly, secure, and sameSite flags
- [ ] **Environment Variables**: Never commit credentials to git
- [ ] **Password Requirements**: Minimum 12 characters, mix of types
- [ ] **Audit Logging**: Log all admin actions (who, what, when)
- [ ] **CSRF Protection**: Add CSRF tokens to all forms
- [ ] **Account Lockout**: Lock account after failed attempts
- [ ] **Password Reset**: Implement secure password reset flow (if Phase 2+)

### Additional Security Headers:
```go
w.Header().Set("X-Frame-Options", "DENY")
w.Header().Set("X-Content-Type-Options", "nosniff")
w.Header().Set("X-XSS-Protection", "1; mode=block")
w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
```

### Recommended Libraries:
- **Hashing**: `golang.org/x/crypto/bcrypt`
- **Sessions**: `github.com/gorilla/sessions`
- **JWT**: `github.com/golang-jwt/jwt/v5`
- **Rate Limiting**: `golang.org/x/time/rate`
- **CSRF**: `github.com/gorilla/csrf`

---

## Migration Path

### Immediate (This Week):
1. Implement Phase 1 (HTTP Basic Auth) - **30 minutes**
2. Deploy to production
3. Test admin access works

### Short Term (Next Week):
1. Plan Phase 2 (Session-Based Auth)
2. Create database migrations
3. Implement login system
4. Test thoroughly
5. Deploy and migrate

### Long Term (Optional):
1. Add OAuth if needed
2. Add audit logging
3. Add admin dashboard
4. Add user management UI

---

## Quick Start: Phase 1 Implementation

If you want to secure the routes NOW, here's the minimal Phase 1 code:

### 1. Create `middleware/auth.go`:
```go
package middleware

import (
    "net/http"
    "os"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()

        adminUser := os.Getenv("ADMIN_USERNAME")
        adminPass := os.Getenv("ADMIN_PASSWORD")

        if !ok || username != adminUser || password != adminPass {
            w.Header().Set("WWW-Authenticate", `Basic realm="Admin Area"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}
```

### 2. Update `main.go`:
```go
import "portfolio-v2/middleware"

// Wrap admin routes
http.HandleFunc("/admin/blog/new", middleware.BasicAuth(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        handlers.NewBlogPageHandler(w, r)
    } else if r.Method == http.MethodPost {
        handlers.CreateBlogPostHandler(db)(w, r)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}))
```

### 3. Set Environment Variables:
```bash
export ADMIN_USERNAME="admin"
export ADMIN_PASSWORD="your-secure-password-here"
```

### 4. Update systemd service:
```ini
[Service]
Environment="ADMIN_USERNAME=admin"
Environment="ADMIN_PASSWORD=your-secure-password-here"
```

**Done!** Admin routes are now password-protected.

---

## Next Steps

1. Choose which phase to implement
2. Review security requirements
3. Create implementation timeline
4. Test in development
5. Deploy to production
6. Monitor and audit

**Question**: Which phase would you like to implement first?
