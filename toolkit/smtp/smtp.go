package smtp

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

const (
	SENDER = "Artotel Group<info@artotelgroup.com>"
	//AUTH_EMAIL    = "agus.sukariyasa@artotelgroup.com"
	AUTH_EMAIL = "artotelgroup@gmail.com"
	//AUTH_PASSWORD = "dinsxjifqrvxiobr"
	AUTH_PASSWORD = "qytfvdfajnlrnigd"
	SMTP_HOST     = "smtp.gmail.com"
	SMTP_PORT     = 587
	DEBUG_CC      = "cahyo@wit.id"
)

func SendMail(ctx context.Context, to []string, subject, message string) error {

	body := "From: " + SENDER + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", SMTP_HOST, SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, AUTH_EMAIL, to, []byte(body))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed send email")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return err
	}

	return nil
}
