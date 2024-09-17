package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthTokenService) ReadAuthToken(ctx context.Context, request sqlc.GetAuthTokenParams) (authToken sqlc.AuthToken, err error) {
	q := sqlc.New(s.mainDB)

	authToken, err = q.GetAuthToken(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get auth token")
		err = errors.WithStack(httpservice.ErrInvalidToken)

		return
	}

	return
}
