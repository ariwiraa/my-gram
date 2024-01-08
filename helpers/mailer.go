package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
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
	message.SetBody("text/html", dm.EmailBody)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

	err := dialer.DialAndSend(message)
	if err != nil {
		log.Println("gagal mengirim email: ", err)
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

	urlString := fmt.Sprintf("%s/verify-email?email=%s&token=%s", os.Getenv("FRONTEND_ORIGIN_URL"), data.Email, data.Token)
	emailBody, _ := h.GenerateHTML(hermes.Email{
		Body: hermes.Body{
			Name: data.Username,
			Intros: []string{
				"Welcome to Lazapedia!",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please click the following button to verify your email. This link expires in 5 minutes.",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  urlString,
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
