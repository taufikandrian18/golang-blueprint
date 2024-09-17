package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/jwt"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthTokenService) AuthToken(ctx context.Context, request payload.AuthTokenPayload) (authToken sqlc.AuthToken, err error) {
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

	// validate app key
	if err = s.validateAppKey(ctx, q, payload.ValidateAppKeyPayload{
		AppName: request.AppName,
		AppKey:  request.AppKey,
	}); err != nil {
		return
	}

	// generate jwt token
	jwtResponse, err := jwt.CreateJWTToken(ctx, s.cfg, jwt.RequestJWTToken{
		AppName:    request.AppName,
		DeviceID:   request.DeviceID,
		DeviceType: request.DeviceType,
		IPAddress:  request.IPAddress,
	})
	if err != nil {
		return
	}

	authToken, err = s.recordToken(ctx, q, jwtResponse, false)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}

func (s *AuthTokenService) RefreshToken(ctx context.Context, request jwt.RequestJWTToken) (authToken sqlc.AuthToken, err error) {
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
				log.FromCtx(ctx).Error(err, "error rollback=%s", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
		}
	}()

	jwtResponse, err := jwt.CreateJWTToken(ctx, s.cfg, request)
	if err != nil {
		return
	}

	authToken, err = s.recordToken(ctx, q, jwtResponse, true)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}

func (s *AuthTokenService) validateAppKey(ctx context.Context, q *sqlc.Queries, request payload.ValidateAppKeyPayload) (err error) {
	appKeyData, err := q.GetAppKeyByName(ctx, request.AppName)
	if err != nil {
		log.FromCtx(ctx).Error(err, "Failed get app key data by name")
		err = errors.WithStack(httpservice.ErrInvalidAppKey)

		return
	}

	if request.AppKey != appKeyData.Key {
		log.FromCtx(ctx).Info("app key is not match")

		err = errors.WithStack(httpservice.ErrInvalidAppKey)

		return
	}

	return
}

func (s *AuthTokenService) recordToken(ctx context.Context, q *sqlc.Queries, token jwt.ResponseJwtToken, isRefreshToken bool) (authToken sqlc.AuthToken, err error) {
	if !isRefreshToken {
		authToken, err = q.InsertAuthToken(ctx, sqlc.InsertAuthTokenParams{
			Name:         token.AppName,
			DeviceID:     token.DeviceID,
			DeviceType:   token.DeviceType,
			Token:        token.Token,
			TokenExpired: token.TokenExpired,
			IpAddress: sql.NullString{
				String: token.IPAddress,
				Valid:  true,
			},
			RefreshToken:        token.RefreshToken,
			RefreshTokenExpired: token.RefreshTokenExpired,
		})
	} else {
		// Get record
		authData, errGetRecord := s.ReadAuthToken(ctx, sqlc.GetAuthTokenParams{
			Name:       token.AppName,
			DeviceID:   token.DeviceID,
			DeviceType: token.DeviceType,
		})
		if errGetRecord != nil {
			err = errGetRecord
			return
		}

		authToken, err = q.InsertAuthToken(ctx, sqlc.InsertAuthTokenParams{
			Name:         authData.Name,
			DeviceID:     authData.DeviceID,
			DeviceType:   authData.DeviceType,
			Token:        token.Token,
			TokenExpired: token.TokenExpired,
			IpAddress: sql.NullString{
				String: token.IPAddress,
				Valid:  true,
			},
			RefreshToken:        token.RefreshToken,
			RefreshTokenExpired: token.RefreshTokenExpired,
			IsLogin:             authData.IsLogin,
			UserLogin:           authData.UserLogin,
		})
	}

	if err != nil {
		log.FromCtx(ctx).Error(err, "failed record token")
		err = errors.WithStack(httpservice.ErrInternalServerError)

		return
	}

	return
}
