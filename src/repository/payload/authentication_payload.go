package payload

import (
	"context"
	"database/sql"
	"time"

	"gitlab.com/wit-id/project-latihan/common/utility"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginPayload struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type RegisterEmployeePayload struct {
	Fullname        string `json:"fullname" valid:"required"`
	Email           string `json:"email" valid:"required"`
	BirthDate       string `json:"birth_date"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password" valid:"required"`
	ConfirmPassword string `json:"confirm_password" valid:"required"`
}

func (payload *LoginPayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

func (payload *RegisterEmployeePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

func ToPayloadToken(accessToken, refreshToken string) (payload tokenResponse) {
	payload = tokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return
}

type LogoutPayload struct {
	AccessToken  string `json:"access_token" valid:"required"`
	RefreshToken string `json:"refresh_token" valid:"required"`
}

func (payload *LogoutPayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

type ChangePassword struct {
	Password             string `json:"password" valid:"required"`
	PasswordConfirmation string `json:"password_confirmation" valid:"required"`
	OldPassword          string `json:"old_password" valid:"required"`
}

func (payload *ChangePassword) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

type ReplacePassword struct {
	Password string `json:"password" valid:"required"`
}

func (payload *ReplacePassword) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

type ForgotPasswordRequestPayload struct {
	Email string `json:"email" valid:"required"`
}

func (payload *ForgotPasswordRequestPayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

type ForgotPasswordSubmitPayload struct {
	Password             string `json:"password" valid:"required"`
	PasswordConfirmation string `json:"password_confirmation" valid:"required"`
}

func (payload *ForgotPasswordSubmitPayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}
	return
}

type readAuthentication struct {
	Guid             string    `json:"guid"`
	EmployeeGuid     string    `json:"employee_guid"`
	EmployeeFullname string    `json:"employee_fullname"`
	Username         string    `json:"username"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
}

func (payload *RegisterEmployeePayload) ToEntity() (data sqlc.Employee) {
	layout := "02-01-2006"

	data = sqlc.Employee{
		Guid:  utility.GenerateGoogleUUID(),
		Email: payload.Email,
	}

	if payload.BirthDate != "" {
		data.DateOfBirth = sql.NullTime{
			Time: utility.ParseStringToTime(
				payload.BirthDate,
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

func ToPayloadAuthentication(data sqlc.GetAuthenticationByIDRow) (response readAuthentication) {
	response = readAuthentication{
		Guid:             data.Guid,
		EmployeeGuid:     data.EmployeeGuid.String,
		EmployeeFullname: data.EmployeeFullname.String,
		Username:         data.Username,
		Status:           data.Status,
		CreatedAt:        data.CreatedAt.Time,
		CreatedBy:        data.CreatedBy,
	}

	return
}
