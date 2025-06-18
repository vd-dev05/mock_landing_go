# Test API endpoints using PowerShell

Write-Host "Testing Go Image API..." -ForegroundColor Green

# Test 1: Root endpoint
Write-Host "`n1. Testing root endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080" -Method GET
    Write-Host "✅ Success: $response" -ForegroundColor Green
} catch {
    Write-Host "❌ Error: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 2: Images endpoint
Write-Host "`n2. Testing images endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "http://localhost:8080/images/" -Method GET
    Write-Host "✅ Success: Status $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "❌ Error: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 3: Non-existent endpoint (should return 404)
Write-Host "`n3. Testing non-existent endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "http://localhost:8080/nonexistent" -Method GET
    Write-Host "✅ Unexpected success: Status $($response.StatusCode)" -ForegroundColor Yellow
} catch {
    Write-Host "✅ Expected 404: $($_.Exception.Message)" -ForegroundColor Green
}

Write-Host "`nAPI testing completed!" -ForegroundColor Green
