package helpers

import (
	"fmt"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

type DataMail struct {
	Username  string
	Email     string
	Token     string
	Code      string
	EmailBody string
	Subject   string
}

func (dm *DataMail) Send() error {
	message := gomail.NewMessage()
	message.SetHeader("From", "MyGram <"+os.Getenv("MAIL_USER")+">")
	message.SetHeader("To", dm.Email)
	message.SetHeader("Subject", dm.Subject)
	message.SetHeader("text/html", dm.EmailBody)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

	err := dialer.DialAndSend(message)
	if err != nil {
		return err
	}

	return nil
}

func Mail(data *DataMail) *DataMail {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "MyGram",
			Link: os.Getenv("FRONTEND_ORIGIN_URL"),
		},
	}

	if data.Code != "" {
		emailBody, _ := h.GenerateHTML(hermes.Email{
			Body: hermes.Body{
				Name: data.Username,
				Intros: []string{
					"Welcome to MyGram!",
				},
				Actions: []hermes.Action{
					{
						Instructions: "Here is your approval code. This code expires in 5 minutes",
						Button: hermes.Button{
							Color: "#22BC66",
							Text:  data.Code,
						},
					},
				},
			},
		})

		return &DataMail{
			Username:  data.Username,
			Email:     data.Email,
			EmailBody: emailBody,
			Subject:   data.Subject,
		}
	}

	urlString := fmt.Sprintf("%s/auth/confirm=%s&token=%s", os.Getenv("FRONTEND_ORIGINAL_URL"), data.Email, data.Token)
	emailBody, _ := h.GenerateHTML(hermes.Email{
		Body: hermes.Body{
			Name: data.Username,
			Intros: []string{
				"Welcome to MyGram!",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Here is your approval code. This code expires in 5 minutes",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  data.Code,
						Link: urlString,
					},
				},
			},
		},
	})

	return &DataMail{
		Username:  data.Username,
		Email:     data.Email,
		EmailBody: emailBody,
		Subject:   data.Subject,
	}
}
