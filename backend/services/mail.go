package services

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
)

// generateRandomCode generates a random verification code.
func generateRandomCode() (string, error) {
	code := make([]byte, 4) // 4 bytes for an 8-character hex string
	if _, err := rand.Read(code); err != nil {
		return "", err
	}
	return hex.EncodeToString(code), nil
}

// SendVerifyCode sends an email with a verification code using SMTP.
func SendVerifyCode(to string) (string, error) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	// Generate a random verification code.
	code, err := generateRandomCode()
	if err != nil {
		return "", err
	}

	// Append the verification code to the email body.
	body := fmt.Sprintf("\n\nYour verification code is: %s", code)

	// Set up the email message.
	msg := []byte("From: " + os.Getenv("MAIL_USER") + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + "验证码" + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Connect to the SMTP server.
	serverName := os.Getenv("MAIL_HOST") + ":" + os.Getenv("MAIL_PORT")
	host, _, _ := net.SplitHostPort(serverName)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", serverName, tlsConfig)
	if err != nil {
		return "", err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return "", err
	}

	// Authenticate.
	if err = client.Auth(auth); err != nil {
		return "", err
	}

	// Set the sender and recipient.
	if err = client.Mail(os.Getenv("MAIL_USER")); err != nil {
		return "", err
	}
	if err = client.Rcpt(to); err != nil {
		return "", err
	}

	// Send the email body.
	w, err := client.Data()
	if err != nil {
		return "", err
	}

	_, err = w.Write(msg)
	if err != nil {
		return "", err
	}

	err = w.Close()
	if err != nil {
		return "", err
	}

	client.Quit()

	log.Println("SendVerifyCode success", to, code)
	return code, nil
}
