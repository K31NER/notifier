# Configuración - Rellena estos datos
$ClientId = "TU_CLIENT_ID_AQUI"         # Ejemplo: "659219...apps.googleusercontent.com"
$ClientSecret = "TU_CLIENT_SECRET_AQUI" # Ejemplo: "GOCSPX-..."
$AuthCode = "TU_AUTH_CODE_AQUI"         # El código que empieza por "4/..." obtenido del navegador

# URL de redirección (debe coincidir exactamente con la usada para obtener el código)
# Si usaste el link manual que te di antes, es esta:
$RedirectUri = "https://developers.google.com/oauthplayground"

# --- No editar debajo de esta línea ---

$params = @{
    client_id     = $ClientId
    client_secret = $ClientSecret
    code          = $AuthCode
    grant_type    = "authorization_code"
    redirect_uri  = $RedirectUri
}

try {
    Write-Host "Intercambiando código por tokens..." -ForegroundColor Cyan
    $response = Invoke-RestMethod -Uri "https://oauth2.googleapis.com/token" -Method Post -Body $params
    
    Write-Host "`n¡ÉXITO!" -ForegroundColor Green
    Write-Host "---------------------------------------------------"
    Write-Host "REFRESH TOKEN (Guarda esto en tu .env):" -ForegroundColor Yellow
    Write-Host $response.refresh_token
    Write-Host "---------------------------------------------------"
    
    if ($response.access_token) {
        Write-Host "Access Token (temporal):" -ForegroundColor Gray
        Write-Host $response.access_token
    }
}
catch {
    Write-Host "`nERROR: No se pudo obtener el token." -ForegroundColor Red
    Write-Host $_.Exception.Message
    Write-Host "Verifica que el código no haya expirado y que redirect_uri coincida." -ForegroundColor DarkRed
}
