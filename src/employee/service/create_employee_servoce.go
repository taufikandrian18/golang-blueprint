package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/utility"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *EmployeeService) CreateEmployee(
	ctx context.Context,
	params sqlc.InsertEmployeeParams) (
	employee sqlc.Employee,
	err error) {

	responseData, err := utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		response, err = query.InsertEmployee(ctx, params)
		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create employee")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	// Get relation data post
	employee, err = s.GetEmployeeByGUID(ctx, responseData.(sqlc.Employee).Guid)
	if err != nil {
		return
	}

	return
}
