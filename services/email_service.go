package services

import (
	"encoding/base64"
	"fmt"
	"github.com/thiago1623/banck_transactions_api/config"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

// EmailService provides methods for sending emails.
type EmailService struct{}

// NewEmailService creates a new EmailService instance.
func NewEmailService() *EmailService {
	return &EmailService{}
}

// SendEmailWithTemplate read the html template for the email
func (es *EmailService) SendEmailWithTemplate(to, subject, templatePath, filePath string) error {
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}
	message := string(templateContent)
	err = es.SendEmailWithAttachment(to, subject, message, filePath)
	if err != nil {
		return err
	}
	return nil
}

// composeMessage Composes the message using the HTML template and the data provided.
func composeMessage(template string, data map[string]string) string {
	for key, value := range data {
		placeholder := "{{" + key + "}}"
		template = strings.Replace(template, placeholder, value, -1)
	}
	return template
}

// SendEmailWithAttachment Send an email with an attachment.
func (es *EmailService) SendEmailWithAttachment(to, subject, body, filePath string) error {
	cfg := config.LoadSettings()
	serverSection := cfg.Section("Server")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	auth := smtp.PlainAuth("", serverSection.Key("SenderEmail").String(),
		serverSection.Key("EmailPassword").String(), "smtp.gmail.com")
	message := ""
	headers := map[string]string{
		"From":         serverSection.Key("SenderEmail").String(),
		"To":           to,
		"Subject":      subject,
		"Content-Type": `multipart/mixed; boundary="BOUNDARY"`,
	}
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n--BOUNDARY\r\n"
	message += "Content-Type: text/html\r\n\r\n" + body + "\r\n"
	message += "--BOUNDARY\r\n"
	message += fmt.Sprintf(`Content-Type: application/octet-stream
Content-Disposition: attachment; filename="%s"
Content-Transfer-Encoding: base64
`, filepath.Base(filePath))
	message += "\r\n"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	message += base64.StdEncoding.EncodeToString(content)
	err = smtp.SendMail("smtp.gmail.com:587",
		auth, serverSection.Key("SenderEmail").String(),
		[]string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
