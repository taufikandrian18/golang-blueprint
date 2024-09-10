package payload

import (
	"context"
	"database/sql"
	"time"

	"gitlab.com/wit-id/project-latihan/common/constants"
	"gitlab.com/wit-id/project-latihan/common/utility"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
)

type InsertEmployeePayload struct {
	Fullname    string `json:"fullname" valid:"required"`
	Email       string `json:"email" valid:"required"`
	PhoneNumber string `json:"phone_number" valid:"required"`
	DateOfBirth string `json:"date_of_birth" valid:"required"`
}

type UpdateEmployeePayload struct {
	Fullname    string `json:"fullname" valid:"required"`
	Email       string `json:"email" valid:"required"`
	PhoneNumber string `json:"phone_number" valid:"required"`
	DateOfBirth string `json:"date_of_birth" valid:"required"`
}

type ListEmployeePayload struct {
	Filter ListFilterEmployeePayload `json:"filter"`
	Limit  int32                     `json:"limit" valid:"required~limit is required field"`
	Offset int32                     `json:"page" valid:"required~page is required field"`
	Order  string                    `json:"order" valid:"required~order is required field"`
	Sort   string                    `json:"sort" valid:"required~sort is required field"` // ASC, DESC
}

type ListFilterEmployeePayload struct {
	SetGuid        bool   `json:"set_guid"`
	Guid           string `json:"guid"`
	SetFullname    bool   `json:"set_fullname"`
	Fullname       string `json:"fullname"`
	SetEmail       bool   `json:"set_email"`
	Email          string `json:"email"`
	SetDateOfBirth bool   `json:"set_date_of_birth"`
	DateOfBirth    string `json:"date_of_birth"`
}

type readEmployeePayload struct {
	GUID        string     `json:"guid,omitempty"`
	EmployeeID  int        `json:"employee_id"`
	Fullname    *string    `json:"fullname"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   *string    `json:"updated_by"`
}

func (payload *InsertEmployeePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	return
}

func (payload *UpdateEmployeePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	return
}

func (payload *ListEmployeePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	return
}

func (payload *InsertEmployeePayload) ToEntity() (data sqlc.InsertEmployeeParams) {

	layout := "02-01-2006"

	data = sqlc.InsertEmployeeParams{
		Guid:      utility.GenerateGoogleUUID(),
		Email:     payload.Email,
		CreatedBy: constants.CreatedByTemporaryBySystem,
	}

	if payload.DateOfBirth != "" {
		data.DateOfBirth = sql.NullTime{
			Time: utility.ParseStringToTime(
				payload.DateOfBirth,
				layout),
			Valid: true,
		}
	}

	if payload.Fullname != "" {
		data.Fullname = sql.NullString{
			String: payload.Fullname,
			Valid:  true,
		}
	}

	if payload.PhoneNumber != "" {
		data.PhoneNumber = sql.NullString{
			String: payload.PhoneNumber,
			Valid:  true,
		}
	}

	return
}

func (payload *UpdateEmployeePayload) ToEntity(key string) (data sqlc.UpdateEmployeeParams) {
	layout := "02-01-2006"
	data = sqlc.UpdateEmployeeParams{
		Email: payload.Email,
		UpdatedBy: sql.NullString{
			String: constants.CreatedByTemporaryBySystem,
			Valid:  true,
		},
		Guid: key,
	}

	if payload.DateOfBirth != "" {
		data.DateOfBirth = sql.NullTime{
			Time: utility.ParseStringToTime(
				payload.DateOfBirth,
				layout),
			Valid: true,
		}
	}

	if payload.Fullname != "" {
		data.Fullname = sql.NullString{
			String: payload.Fullname,
			Valid:  true,
		}
	}

	if payload.PhoneNumber != "" {
		data.PhoneNumber = sql.NullString{
			String: payload.PhoneNumber,
			Valid:  true,
		}
	}

	return
}

func (payload *ListEmployeePayload) ToEntity() (data sqlc.ListEmployeeParams) {
	layout := "02-01-2006"
	dateString := utility.ParseStringToTime(
		payload.Filter.DateOfBirth,
		layout)
	data = sqlc.ListEmployeeParams{
		SetGuid:        payload.Filter.SetGuid,
		Guid:           payload.Filter.Guid,
		SetFullname:    payload.Filter.SetFullname,
		Fullname:       queryStringLike(payload.Filter.Fullname),
		SetEmail:       payload.Filter.SetEmail,
		Email:          queryStringLike(payload.Filter.Email),
		SetDateOfBirth: payload.Filter.SetDateOfBirth,
		DateOfBirth:    dateString,
		OrderParam:     makeOrderParam(payload.Order, payload.Sort),
		OffsetPage:     makeOffset(payload.Limit, payload.Offset),
		LimitData:      limitWithDefault(payload.Limit),
	}

	return
}

func ToPayloadEmployee(data sqlc.Employee) (payload readEmployeePayload) {
	payload = readEmployeePayload{
		GUID:       data.Guid,
		EmployeeID: int(data.EmployeeID.Int32),
		Email:      data.Email,
		CreatedAt:  data.CreatedAt.Time,
		CreatedBy:  data.CreatedBy,
	}

	if data.Fullname.Valid {
		payload.Fullname = &data.Fullname.String
	}

	if data.PhoneNumber.Valid {
		payload.PhoneNumber = &data.PhoneNumber.String
	}

	if data.DateOfBirth.Valid {
		payload.DateOfBirth = &data.DateOfBirth.Time
	}

	if data.UpdatedAt.Valid {
		payload.UpdatedAt = &data.UpdatedAt.Time
	}

	if data.UpdatedBy.Valid {
		payload.UpdatedBy = &data.UpdatedBy.String
	}

	return
}

func ToPayloadGetEmployee(data sqlc.Employee) (payload readEmployeePayload) {
	payload = readEmployeePayload{
		GUID:       data.Guid,
		EmployeeID: int(data.EmployeeID.Int32),
		Email:      data.Email,
		CreatedAt:  data.CreatedAt.Time,
		CreatedBy:  data.CreatedBy,
	}

	if data.Fullname.Valid {
		payload.Fullname = &data.Fullname.String
	}

	if data.PhoneNumber.Valid {
		payload.PhoneNumber = &data.PhoneNumber.String
	}

	if data.DateOfBirth.Valid {
		payload.DateOfBirth = &data.DateOfBirth.Time
	}

	if data.UpdatedAt.Valid {
		payload.UpdatedAt = &data.UpdatedAt.Time
	}

	if data.UpdatedBy.Valid {
		payload.UpdatedBy = &data.UpdatedBy.String
	}

	return
}

func ToPayloadListEmployee(collection []sqlc.Employee) (payload []*readEmployeePayload) {
	payload = make([]*readEmployeePayload, len(collection))

	for i := range collection {
		payload[i] = new(readEmployeePayload)
		data := readEmployeePayload{
			GUID:       collection[i].Guid,
			EmployeeID: int(collection[i].EmployeeID.Int32),
			Email:      collection[i].Email,
			CreatedAt:  collection[i].CreatedAt.Time,
			CreatedBy:  collection[i].CreatedBy,
		}

		if collection[i].Fullname.Valid {
			data.Fullname = &collection[i].Fullname.String
		}

		if collection[i].PhoneNumber.Valid {
			data.PhoneNumber = &collection[i].PhoneNumber.String
		}

		if collection[i].DateOfBirth.Valid {
			data.DateOfBirth = &collection[i].DateOfBirth.Time
		}

		if collection[i].UpdatedAt.Valid {
			data.UpdatedAt = &collection[i].UpdatedAt.Time
		}

		if collection[i].UpdatedBy.Valid {
			data.UpdatedBy = &collection[i].UpdatedBy.String
		}

		payload[i] = &data
	}

	return
}

func ToEntityDeleteEmployee(key, actorID string) (data sqlc.UpdateEmployeeStatusParams) {
	data = sqlc.UpdateEmployeeStatusParams{
		Status:    constants.StatusDeleted,
		UpdatedBy: sql.NullString{String: actorID, Valid: true},
		Guid:      key,
	}

	return
}
