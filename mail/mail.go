package mail

import (
	"fmt"
	"time"

	"github.com/delphinus/moesia/config"
	"gopkg.in/gomail.v2"
)

// Mail is a struct for sending mail
type Mail struct {
	config *config.Config
}

// New returns a new instance for Mail
func New(config *config.Config) *Mail {
	return &Mail{config}
}

// Send sends the mail
func (m *Mail) Send(body string) error {
	mes := gomail.NewMessage()
	mes.SetHeader("From", m.config.From)
	mes.SetHeader("To", m.config.To...)
	mes.SetHeader("Cc", m.config.Cc...)
	mes.SetHeader("Subject", m.title())
	mes.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, m.config.GmailUserName, m.config.GmailPassword)
	if err := d.DialAndSend(mes); err != nil {
		return fmt.Errorf("failed to DialAndSend: %v", err)
	}
	return nil

}

func (m *Mail) title() string {
	return time.Now().Format("ITS 宿泊施設空き情報（2006-01-02 15:04:05）")
}
