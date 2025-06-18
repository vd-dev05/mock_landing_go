# Production mode
$env:ENV_MODE="production"
Write-Host "Starting server in production mode..."
go run main.go
