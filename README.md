# Go Image API

A simple REST API for serving images built with Go.

## Features

- 🚀 Static file serving for images
- 📊 Health check endpoint
- 🔗 JSON API for image metadata
- 🌍 Environment-based configuration
- ⚡ Lightweight and fast

## Endpoints

### Root
```
GET /
```
Returns API status and version.

### Health Check
```
GET /health
```
Returns health status with timestamp.

### Image List API
```
GET /api/images
```
Returns JSON list of available images with metadata.

### Static Images
```
GET /images/{filename}
```
Serves static image files.

## Development

### Prerequisites
- Go 1.21+
- Git

### Setup
```bash
# Clone repository
git clone <repository-url>
cd golang

# Install dependencies
go mod download

# Run in development mode
.\dev.ps1
# or
make dev
```

### Available Scripts
- `.\dev.ps1` - Run in development mode
- `.\prod.ps1` - Run in production mode
- `.\build.ps1` - Build binary
- `make dev` - Development with Makefile
- `make prod` - Production with Makefile
- `make build` - Build with Makefile

## Environment Variables

- `PORT` - Server port (default: 8080)
- `ENV_MODE` - Environment mode (development/production)
- `BASE_URL` - Base URL for image links

## Deployment

### Railway
1. Push to GitHub
2. Connect Railway to repository
3. Set environment variables
4. Deploy automatically

## API Testing

Use the provided `api-test.http` file with REST Client extension in VS Code.

## License

MIT
