# Favicon Setup Instructions

Your website is configured to use favicons, but you need to create the actual favicon image files.

## Required Files

Place these files in the `/static` directory:

1. **favicon-32x32.png** - 32x32 pixel PNG (standard favicon)
2. **favicon-16x16.png** - 16x16 pixel PNG (small favicon)
3. **apple-touch-icon.png** - 180x180 pixel PNG (iOS home screen icon)

## Option 1: Use a Favicon Generator (Easiest)

### Recommended: Favicon.io
1. Visit https://favicon.io/
2. Choose one of these options:
   - **Text** - Generate from your initials (e.g., "MH")
   - **Image** - Upload a logo/image you have
   - **Emoji** - Use an emoji as your favicon

3. Customize:
   - Font: Choose a clean, professional font
   - Colors: Use your brand colors (purple/blue: #8b5cf6, #3b82f6)
   - Background: Black (#000000) to match your site theme

4. Download the generated package
5. Extract and copy these files to `/static`:
   - `favicon-32x32.png`
   - `favicon-16x16.png`
   - `apple-touch-icon.png`

### Alternative: RealFaviconGenerator
1. Visit https://realfavicongenerator.net/
2. Upload your image/logo
3. Customize for different platforms
4. Download and extract to `/static`

## Option 2: Design Your Own

If you have design software (Figma, Photoshop, etc.):

1. Create a square canvas (at least 512x512px)
2. Design your icon/logo
3. Export at these sizes:
   - 32x32px → `favicon-32x32.png`
   - 16x16px → `favicon-16x16.png`
   - 180x180px → `apple-touch-icon.png`

### Design Tips
- Keep it simple - favicons are small
- Use high contrast
- Avoid fine details
- Test at actual size (16px is tiny!)
- Consider using your initials or a geometric shape

## Option 3: Use an Emoji

Quick temporary solution using an emoji:

1. Visit https://favicon.io/emoji-favicons/
2. Search for an emoji (💻 📱 ⚡ 🚀 etc.)
3. Download the package
4. Copy files to `/static`

## Option 4: Simple Text-Based Icon

Create a simple colored square with your initials:

1. Visit https://favicon.io/favicon-generator/
2. Text: **MH**
3. Background: **Rounded** with **#8b5cf6** (purple)
4. Font: **Roboto** or **Inter**
5. Font Color: **#ffffff** (white)
6. Download and copy to `/static`

## After Adding Favicon Files

1. Generate templates: `templ generate`
2. Rebuild: `go build -o portfolio-v2 main.go`
3. Test locally: `./portfolio-v2` then visit http://localhost:8080
4. Check your browser tab - you should see the favicon!
5. Deploy: `./deploy/full-deploy.sh`

## Testing

- **Browser tab**: Should show your favicon
- **Bookmarks**: Should display correctly when bookmarked
- **iOS home screen**: Add to home screen to test apple-touch-icon
- **Hard refresh**: Press Ctrl+F5 (or Cmd+Shift+R on Mac) to clear cache

## Current Status

✅ HTML is configured to load favicons
✅ site.webmanifest is created
❌ Actual favicon image files need to be created

Once you add the PNG files, your favicon will appear automatically!
