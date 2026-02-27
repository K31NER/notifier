package models

type EmailType string

const (
	EmailTypeInfo    EmailType = "INFO"
	EmailTypeWarning EmailType = "WARNING"
	EmailTypeError   EmailType = "ERROR"
)

type Body struct {
	CompanyLogo string `json:"company_logo"`
	Message     string `json:"message" binding:"required"`
}

type EmailRequest struct {
	Recipient   string    `json:"recipient" binding:"required,email"`
	Type        EmailType `json:"type" binding:"required,oneof=INFO WARNING ERROR"`
	CompanyName string    `json:"company_name" binding:"required"`
	Subject     string    `json:"subject" binding:"required"`
	Body        Body      `json:"body" binding:"required"`
}