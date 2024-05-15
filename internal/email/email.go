package email

import (
	"fmt"
	"log"
	"net/smtp"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
)

type EmailService struct {
	Env *bootstrap.Env
}

type EmailServiceInterface interface {
	SendEmail(requestEmail dto.EmailRequest) error
}

func NewEmailService() EmailServiceInterface {
	return &EmailService{
		Env: bootstrap.NewEnv(),
	}
}

func buildMessage(from string, requestEmail dto.EmailRequest) string {
	message := fmt.Sprintf("From: %s\r\n", from)
	message += fmt.Sprintf("To: %s\r\n", requestEmail.OrganizationEmail)
	message += fmt.Sprintf("Subject: %s\r\n", requestEmail.ProposalSubject)
	message += "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/html; charset=UTF-8\r\n"
	message += "\r\n"
	message += fmt.Sprintf("<p><strong>Nama Pengirim:</strong> %s</p>", requestEmail.UserFullName)
	message += fmt.Sprintf("<p><strong>Judul Proposal:</strong> %s</p>", requestEmail.ProposalSubject)
	message += fmt.Sprintf("<p><strong>Deskripsi Proposal:</strong> %s</p>", requestEmail.ProposalContent)
	message += fmt.Sprintf("<p><strong>Lampiran Proposal:</strong> <a href=\"%s\">Download Proposal</a></p>", requestEmail.ProposalAttachment)
	return message
}

func (es *EmailService) SendEmail(requestEmail dto.EmailRequest) (err error) {

	auth := smtp.PlainAuth("", es.Env.SMTP_USERNAME, es.Env.SMTP_PASSWORD, es.Env.SMTP_SERVER)

	msg := buildMessage(es.Env.SMTP_USERNAME, requestEmail)

	err = smtp.SendMail(es.Env.SMTP_SERVER+":"+es.Env.SMTP_PORT, auth, es.Env.SMTP_USERNAME, []string{requestEmail.OrganizationEmail}, []byte(msg))

	if err != nil {
		log.Println("error send email: ", err)
		log.Fatal(err)
		return err
	}

	return nil
}
