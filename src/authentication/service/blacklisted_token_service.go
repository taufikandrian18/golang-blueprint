package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthenticationService) IsTokenBlacklisted(ctx context.Context, token string) (
	valid bool, err error) {
	query := sqlc.New(s.mainDB)

	blacklistedToken, err := query.GetBlacklistedToken(ctx, sql.NullString{
		String: token,
		Valid:  true})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get blacklisted token")
		err = errors.WithStack(httpservice.ErrDataNotFound)

		return
	}

	if blacklistedToken.Token.Valid {
		valid = true
	}

	return
}

func (s *AuthenticationService) Logout(ctx context.Context,
	logout payload.LogoutPayload) (
	valid bool, err error) {

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

	_, err = q.InsertBlacklistedToken(ctx, sqlc.InsertBlacklistedTokenParams{
		Token: sql.NullString{String: logout.AccessToken, Valid: true},
		Type:  sql.NullString{String: "access_token", Valid: true},
	})

	_, err = q.InsertBlacklistedToken(ctx, sqlc.InsertBlacklistedTokenParams{
		Token: sql.NullString{String: logout.RefreshToken, Valid: true},
		Type:  sql.NullString{String: "refresh_token", Valid: true},
	})

	if err == nil {
		valid = true
	}

	return
}
