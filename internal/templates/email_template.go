package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/K31NER/notifier.git/internal/models"
)

const EmailTemplateHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f9fafb; margin: 0; padding: 20px; }
        .container { background-color: #ffffff; padding: 30px; border-radius: 8px; max-width: 600px; margin: 0 auto; box-shadow: 0 4px 6px rgba(0,0,0,0.05); }
        .header { text-align: center; border-bottom: 2px solid #f3f4f6; padding-bottom: 20px; margin-bottom: 20px; }
        .logo { max-width: 150px; max-height: 80px; object-fit: contain; }
        .company-name { font-size: 24px; font-weight: 600; color: #1f2937; margin-top: 15px; }
        .content { font-size: 16px; color: #4b5563; line-height: 1.6; white-space: pre-wrap; }
        .footer { margin-top: 40px; text-align: center; font-size: 12px; color: #9ca3af; border-top: 1px solid #f3f4f6; padding-top: 20px; }
        .badge { display: inline-block; padding: 4px 8px; border-radius: 4px; font-size: 12px; font-weight: bold; color: white; margin-bottom: 10px; }
        .badge-INFO { background-color: #3b82f6; }
        .badge-WARNING { background-color: #f59e0b; }
        .badge-ERROR { background-color: #ef4444; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            {{if .CompanyLogo}}
            <img src="{{.CompanyLogo}}" alt="{{.CompanyName}} Logo" class="logo">
            {{end}}
            <div class="company-name">{{.CompanyName}}</div>
            <div class="badge badge-{{.Type}}">{{.Type}}</div>
        </div>
        <div class="content">{{.Message}}</div>
        <div class="footer">
            &copy; {{.Year}} {{.CompanyName}}. Todos los derechos reservados.
        </div>
    </div>
</body>
</html>
`

// BuildEmailMessage genera el string completo del correo en formato MIME HTML
func BuildEmailMessage(email models.EmailRequest) (string, error) {
	// 1. Formatear el asunto
	subject := fmt.Sprintf("[%s] %s - %s", email.Type, email.CompanyName, email.Subject)

	// 2. Parsear la plantilla HTML
	tmpl, err := template.New("email").Parse(EmailTemplateHTML)
	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}

	// 3. Ejecutar la plantilla con los datos
	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, map[string]interface{}{
		"CompanyLogo": email.Body.CompanyLogo,
		"CompanyName": email.CompanyName,
		"Type":        email.Type,
		"Message":     email.Body.Message,
		"Year":        time.Now().Year(),
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute email template: %w", err)
	}

	// 4. Construir el mensaje MIME (HTML)
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msgStr := fmt.Sprintf("To: %s\r\nSubject: %s\r\n%s%s", email.Recipient, subject, mimeHeaders, bodyBuffer.String())

	return msgStr, nil
}