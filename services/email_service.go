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

// EmailService proporciona métodos para enviar correos electrónicos.
type EmailService struct{}

// NewEmailService crea una nueva instancia de EmailService.
func NewEmailService() *EmailService {
	return &EmailService{}
}

func (es *EmailService) SendEmailWithTemplate(to, subject, templatePath, filePath string) error {
	// Leer el contenido del archivo de plantilla HTML
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Componer el mensaje utilizando la plantilla
	message := string(templateContent)

	// Enviar el correo electrónico
	err = es.SendEmailWithAttachment(to, subject, message, filePath)
	if err != nil {
		return err
	}

	return nil
}

// composeMessage compone el mensaje utilizando la plantilla HTML y los datos proporcionados.
func composeMessage(template string, data map[string]string) string {
	for key, value := range data {
		placeholder := "{{" + key + "}}"
		template = strings.Replace(template, placeholder, value, -1)
	}
	return template
}

// SendEmailWithAttachment envía un correo electrónico con un archivo adjunto.
func (es *EmailService) SendEmailWithAttachment(to, subject, body, filePath string) error {
	cfg := config.LoadSettings()
	serverSection := cfg.Section("Server")
	// Abrir el archivo a adjuntar
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Configurar el cliente SMTP
	auth := smtp.PlainAuth("", serverSection.Key("SenderEmail").String(),
		serverSection.Key("EmailPassword").String(), "smtp.gmail.com")

	// Componer el mensaje
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

	// Leer y codificar el contenido del archivo adjunto
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	message += base64.StdEncoding.EncodeToString(content)

	// Enviar el correo electrónico
	err = smtp.SendMail("smtp.gmail.com:587",
		auth, serverSection.Key("SenderEmail").String(),
		[]string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
