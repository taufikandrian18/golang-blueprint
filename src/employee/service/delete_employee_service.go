package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/constants"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/utility"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *EmployeeService) DeleteEmployee(ctx context.Context, guid string) (err error) {
	data := sqlc.UpdateEmployeeStatusParams{
		Status: constants.StatusDeleted,
		UpdatedBy: sql.NullString{
			String: constants.CreatedByTemporaryBySystem,
			Valid:  true,
		},
		Guid: guid,
	}
	_, err = utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		response, err = query.UpdateEmployeeStatus(ctx, data)

		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed delete employee")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
