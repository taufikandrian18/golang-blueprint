package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/utility"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
	"gitlab.com/wit-id/project-latihan/toolkit/smtp"
)

func (s *AuthenticationService) ChangeEmployeePassword(ctx context.Context,
	username string, params payload.ChangePassword) (
	err error) {

	query := sqlc.New(s.mainDB)
	u, err := query.GetAuthenticationByUsername(ctx, username)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get User")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	oldHashedPassword := utility.HashPassword(params.OldPassword, u.Salt.String)
	if oldHashedPassword != u.Password {
		log.FromCtx(ctx).Error(err, "old password do not match")
		err = errors.WithStack(httpservice.ErrPasswordNotMatch)
		return
	}

	if params.Password != params.PasswordConfirmation {
		log.FromCtx(ctx).Error(err, "new password do not match")
		err = errors.WithStack(httpservice.ErrPasswordNotMatch)
		return
	}

	newHashedPassword := utility.HashPassword(params.Password, u.Salt.String)

	err = query.UpdateAuthenticationPassword(ctx, sqlc.UpdateAuthenticationPasswordParams{
		Password:  newHashedPassword,
		UpdatedBy: u.EmployeeGuid,
		Guid:      u.Guid,
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update password")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	return
}

func (s *AuthenticationService) ForgotPasswordRequest(ctx context.Context, username string) (err error) {
	q := sqlc.New(s.mainDB)
	u, _ := q.GetAuthenticationByUsername(ctx, username)
	if u.Guid == "" {
		log.FromCtx(ctx).Error(err, "user not found")
		err = errors.WithStack(httpservice.ErrUserNotFound)
		return
	}

	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		log.FromCtx(ctx).Error(err, "user not found")
		err = errors.WithStack(httpservice.ErrUserNotFound)
		return
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	err = q.UpdateAuthenticationForgotPassword(ctx, sqlc.UpdateAuthenticationForgotPasswordParams{
		Guid:                 u.Guid,
		ForgotPasswordToken:  sql.NullString{Valid: true, String: token},
		ForgotPasswordExpiry: sql.NullTime{Valid: true, Time: time.Now().Add(15 * time.Minute)},
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update forgot token")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	resetLink := fmt.Sprint(s.cfg.GetString("reset-password.url"), token)
	// from := mail.NewEmail("Cahyo Febrianto", "cahyo@wit.id")
	subject := "Artotel Forgot Password"
	// to := mail.NewEmail(u.EmployeeFullname.String, u.Username)
	plainTextContent := fmt.Sprintf("Hello,\n\nYou have requested to reset your password. Please click on the link below to reset your password:\n\n%s\n\nIf you didn't initiate this request, you can safely ignore this email.\n\nThanks,\nThe Artotel Team", resetLink)
	// htmlContent := fmt.Sprintf("Hello,\n\nYou have requested to reset your password. Please click on the link below to reset your password:\n\n%s\n\nIf you didn't initiate this request, you can safely ignore this email.\n\nThanks,\nThe Artotel Team", resetLink)
	// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// client := sendgrid.NewSendClient("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")
	// _, err = client.Send(message)
	// if err != nil {
	// 	log.FromCtx(ctx).Error(err, "user failed to send mail")
	// 	err = errors.WithStack(httpservice.ErrUnknownSource)
	// 	return
	// }

	err = smtp.SendMail(ctx, []string{u.Username}, subject, plainTextContent)
	if err != nil {
		return
	}

	return
}

func (s *AuthenticationService) ForgotPasswordChange(ctx context.Context, token, password string) (err error) {
	q := sqlc.New(s.mainDB)
	u, _ := q.GetAuthenticationByForgotPasswordToken(ctx, sql.NullString{Valid: true, String: token})
	if u.Guid == "" {
		log.FromCtx(ctx).Error(err, "user not found")
		err = errors.WithStack(httpservice.ErrUserNotFound)
		return
	}

	if time.Now().After(u.ForgotPasswordExpiry.Time) {
		log.FromCtx(ctx).Error(err, "expired token")
		err = errors.WithStack(httpservice.ErrVoucherIsExpired)
		return
	}

	encryptedPassword := utility.HashPassword(password, u.Salt.String)
	err = q.UpdateAuthenticationPassword(ctx, sqlc.UpdateAuthenticationPasswordParams{
		Password:  encryptedPassword,
		UpdatedBy: u.EmployeeGuid,
		Guid:      u.Guid,
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update password")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	return
}
