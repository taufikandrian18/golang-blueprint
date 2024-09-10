package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *EmployeeService) GetEmployeeByGUID(ctx context.Context, employeeGUID string) (
	employee sqlc.Employee, err error) {
	query := sqlc.New(s.mainDB)

	employee, err = query.GetEmployee(ctx, employeeGUID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get employee")
		err = errors.WithStack(httpservice.ErrDataNotFound)

		return
	}

	return
}

func (s *EmployeeService) ListEmployee(ctx context.Context, request sqlc.ListEmployeeParams) (
	listEmployee []sqlc.Employee, totalData int64, err error) {
	query := sqlc.New(s.mainDB)

	// Get Total Data
	totalData, err = s.getTotalDataEmployee(ctx, query, request)
	if err != nil {
		return
	}

	listEmployee, err = query.ListEmployee(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed read list employee")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}

func (s *EmployeeService) getTotalDataEmployee(
	ctx context.Context,
	query *sqlc.Queries,
	request sqlc.ListEmployeeParams) (totalData int64, err error) {
	requestParam := sqlc.CountEmployeeParams{
		SetGuid:        request.SetGuid,
		Guid:           request.Guid,
		SetFullname:    request.SetFullname,
		Fullname:       request.Fullname,
		SetEmail:       request.SetEmail,
		Email:          request.Email,
		SetDateOfBirth: request.SetDateOfBirth,
		DateOfBirth:    request.DateOfBirth,
	}

	totalData, err = query.CountEmployee(ctx, requestParam)
	if err != nil {
		log.FromCtx(ctx).Error(err, "Failed get total data employee")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
