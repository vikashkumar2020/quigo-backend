package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/k3a/html2text"
	"github.com/vikashkumar2020/quigo-backend/app/models"
	config "github.com/vikashkumar2020/quigo-backend/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, data *EmailData, temp string) {
	config := config.NewEmailConfig()

	// Sender data.
	from := config.EmailFrom
	smtpPass := config.EmailPass
	smtpUser := config.EmailUser
	to := user.Email
	smtpHost := config.EmailHost
	smtpPort := config.EmailPort

	var body bytes.Buffer

	template, err := ParseTemplateDir("utils/template")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, temp, data)

	m := gomail.NewMessage()
	fmt.Println(data.URL)
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Println(err)
	}
	d := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}
