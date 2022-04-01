package mailer

import (
	"fmt"
	"net/smtp"
)

const (
	// gmail host
	host = "smtp.gmail.com"
	// default port of smtp server
	port = "587"
)

type Client struct {
	from     string
	password string
}

func NewClient(from, password string) *Client {
	return &Client{from: from, password: password}
}

const emailTemplate = "From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s"

func (c *Client) Send(to string, body string) error {
	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", c.from, c.password, host)

	toList := []string{to}
	return smtp.SendMail(host+":"+port, auth, c.from, toList, []byte(fmt.Sprintf(emailTemplate, c.from, to, "Crypto Arbitrage - Jackpot", body)))
}
