package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/constants"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/utility"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func (s *AuthenticationService) RegisterEmployee(ctx context.Context,
	emp sqlc.Employee, request payload.RegisterEmployeePayload) (
	err error) {

	paramInsers := sqlc.InsertEmployeeParams{
		Guid:        emp.Guid,
		Fullname:    emp.Fullname,
		Email:       emp.Email,
		PhoneNumber: emp.PhoneNumber,
		DateOfBirth: emp.DateOfBirth,
		Status:      constants.StatusActive,
		CreatedBy:   constants.CreatedByTemporaryBySystem,
	}

	responseData, err := utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		response, err = query.InsertEmployee(ctx, paramInsers)
		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create employee")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	salt := utility.GenerateSalt()
	encryptedPassword := utility.HashPassword(request.ConfirmPassword, salt)

	params := sqlc.InsertAuthenticationParams{
		Guid: utility.GenerateGoogleUUID(),
		EmployeeGuid: sql.NullString{
			String: responseData.(sqlc.Employee).Guid,
			Valid:  true},
		Username:  responseData.(sqlc.Employee).Email,
		Password:  encryptedPassword,
		Status:    constants.StatusActive,
		CreatedBy: emp.CreatedBy,
		Salt:      sql.NullString{String: salt, Valid: true},
	}

	query := sqlc.New(s.mainDB)

	_, err = query.InsertAuthentication(ctx, params)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed insert authentication")
		err = errors.WithStack(err)
		return
	}

	return
}
