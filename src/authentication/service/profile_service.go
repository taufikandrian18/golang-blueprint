package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthenticationService) GetProfileFromToken(ctx context.Context, guid string) (
	emp sqlc.Employee, err error) {
	query := sqlc.New(s.mainDB)

	emp, err = query.GetEmployee(ctx, guid)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get Employee")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	return
}

func (s *AuthenticationService) GetAuthenticationByUsername(ctx context.Context, username string) (
	emp sqlc.GetAuthenticationByUsernameRow, err error) {
	query := sqlc.New(s.mainDB)

	emp, err = query.GetAuthenticationByUsername(ctx, username)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get Employee")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	return
}
