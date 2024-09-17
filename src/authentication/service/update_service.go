package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/utility"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthenticationService) UpdateAuthenticationUsernameByEmployeeID(ctx context.Context, request sqlc.UpdateAuthenticationUsernameByEmployeeIDParams) (err error) {
	_, err = utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		err = query.UpdateAuthenticationUsernameByEmployeeID(ctx, request)
		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed update Employee")
		err = errors.WithStack(utility.ParseError(err))
		return
	}

	return
}
