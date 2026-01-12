# Edit Routes Implementation Summary

## Overview

Added PUT routes and edit templates for both blogs and projects, allowing authenticated users to update existing content.

## Files Created

### Templates
1. **templates/edit-blog.templ** - Edit form for blog posts
   - Pre-filled form with existing blog post data
   - Reuses new-blog.css styling
   - Includes success page after update

2. **templates/edit-project.templ** - Edit form for projects
   - Pre-filled form with existing project data
   - Reuses new-project.css styling
   - Includes success page after update

### Handlers
3. **handlers/edit-blog.go**
   - `EditBlogPageHandler` - GET /admin/blog/{id} - Shows edit form
   - `UpdateBlogHandler` - POST /admin/blog/{id} - Processes updates

4. **handlers/edit-project.go**
   - `EditProjectPageHandler` - GET /admin/project/{id} - Shows edit form
   - `UpdateProjectHandler` - POST /admin/project/{id} - Processes updates

## Files Modified

### Database Functions

5. **database/blog.go**
   - Added `UpdateBlogPost(db, id, title, excerpt, content, tags)` - Updates blog post
   - Added `GetBlogPostByID(db, id)` - Fetches blog by ID for editing

6. **database/project.go**
   - Added `UpdateProject(db, project)` - Updates project
   - Added `GetProjectByID(db, id)` - Fetches project by ID for editing

### Routes

7. **main.go**
   - Added `/admin/blog/{id}` - GET (edit form) and POST (update)
   - Added `/admin/project/{id}` - GET (edit form) and POST (update)
   - Both routes protected with HTTP Basic Auth

## Routes Summary

### Blog Routes
| Method | URL | Description | Auth |
|--------|-----|-------------|------|
| GET | `/admin/blog/new` | New blog form | ✓ |
| POST | `/admin/blog/new` | Create blog | ✓ |
| GET | `/admin/blog/{id}` | Edit blog form | ✓ |
| POST | `/admin/blog/{id}` | Update blog (PUT) | ✓ |

### Project Routes
| Method | URL | Description | Auth |
|--------|-----|-------------|------|
| GET | `/admin/project/new` | New project form | ✓ |
| POST | `/admin/project/new` | Create project | ✓ |
| GET | `/admin/project/{id}` | Edit project form | ✓ |
| POST | `/admin/project/{id}` | Update project (PUT) | ✓ |

## Implementation Details

### PUT Method Override
Since HTML forms only support GET and POST, we use a hidden field to simulate PUT:
```html
<input type="hidden" name="_method" value="PUT"/>
```

The handlers check for this field:
```go
if r.FormValue("_method") != "PUT" {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
```

### URL Structure
- Edit blog: `/admin/blog/{id}` (e.g., `/admin/blog/1`)
- Edit project: `/admin/project/{id}` (e.g., `/admin/project/2`)

The ID is extracted from the URL path:
```go
pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
id, err := strconv.Atoi(pathParts[2])
```

### Data Preservation
- **Blog Posts**: Title, excerpt, content, and tags are pre-filled
- **Projects**: Title, description, technologies, GitHub URL, image URL, and featured status are pre-filled
- **Slugs**: Preserved (not editable) to maintain URL consistency

### Validation
- Required fields validated
- Same validation as create forms
- Returns 400 Bad Request for missing fields
- Returns 404 Not Found if blog/project doesn't exist

## Usage Examples

### Edit a Blog Post
1. Navigate to: `https://michaelhegner.com/admin/blog/1`
2. Browser prompts for admin credentials
3. Edit form appears with existing data
4. Make changes and click "Update Post"
5. Success message displayed

### Edit a Project
1. Navigate to: `https://michaelhegner.com/admin/project/1`
2. Browser prompts for admin credentials
3. Edit form appears with existing data
4. Make changes and click "Update Project"
5. Success message displayed

## Security

- All edit routes protected with HTTP Basic Auth middleware
- Same authentication as create routes
- Username and password required (set during deployment)
- HTTPS encryption via Caddy in production

## Testing Locally

```bash
# Set credentials
export ADMIN_USERNAME="admin"
export ADMIN_PASSWORD="testpass12345"

# Run server
go run main.go

# Visit edit pages:
# http://localhost:8080/admin/blog/1
# http://localhost:8080/admin/project/1
```

## Next Steps (Optional)

- Add delete functionality (DELETE method)
- Add list/manage pages showing all blogs/projects with edit links
- Add image upload functionality
- Add draft/publish workflow
- Add version history/audit trail

---

**Status**: ✅ Ready to use! All edit routes are functional and protected.
