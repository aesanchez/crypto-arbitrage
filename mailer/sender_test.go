package mailer

import "testing"

func TestSend(t *testing.T) {
	email := ""
	pass := ""
	to := ""
	c := NewClient(email, pass)
	if err := c.Send(to, "Hello World!"); err != nil {
		t.Fatal(err)
	}
}
