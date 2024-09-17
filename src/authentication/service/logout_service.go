package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/jwt"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthenticationService) LogoutToken(ctx context.Context, request jwt.RequestJWTToken) (err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
		}
	}()

	if err = q.ClearAuthTokenUserLogin(ctx, sqlc.ClearAuthTokenUserLoginParams{
		Name:       request.AppName,
		DeviceID:   request.DeviceID,
		DeviceType: request.DeviceType,
	}); err != nil {
		log.FromCtx(ctx).Error(err, "failed clear auth user login")
		err = errors.WithStack(err)

		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
