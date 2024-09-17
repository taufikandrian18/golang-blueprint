package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/jwt"
	"gitlab.com/wit-id/project-latihan/common/utility"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthenticationService) Login(ctx context.Context, request payload.LoginPayload, jwtRequest jwt.RequestJWTToken) (
	u sqlc.GetAuthenticationByUsernameRow, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	query := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
		}
	}()

	u, err = query.GetAuthenticationByUsername(ctx, request.Username)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get User")
		err = errors.WithStack(httpservice.ErrUserNotMatch)
		return
	}

	hashedPassword := utility.HashPassword(request.Password, u.Salt.String)
	if hashedPassword != u.Password {
		log.FromCtx(ctx).Error(err, "password do not match")
		err = errors.WithStack(httpservice.ErrPasswordNotMatch)
		return
	}

	// Update Last login user backoffice
	if err = query.RecordAuthenticationLastLogin(ctx, u.Guid); err != nil {
		log.FromCtx(ctx).Error(err, "failed record last login")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	// Update token auth record
	if err = query.RecordAuthTokenUserLogin(ctx, sqlc.RecordAuthTokenUserLoginParams{
		UserLogin: sql.NullString{
			String: u.Guid,
			Valid:  true,
		},
		Name:       jwtRequest.AppName,
		DeviceID:   jwtRequest.DeviceID,
		DeviceType: jwtRequest.DeviceType,
	}); err != nil {
		log.FromCtx(ctx).Error(err, "failed update token auth login user")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
