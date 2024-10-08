package payload

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
)

// AuthTokenPayload ...
type AuthTokenPayload struct {
	AppName    string `json:"app_name" valid:"required"`
	AppKey     string `json:"app_key" valid:"required"`
	DeviceID   string `json:"device_id" valid:"required"`
	DeviceType string `json:"device_type" valid:"required"`
	IPAddress  string `json:"ip_address" valid:"required"`
}

type readAuthTokenPayload struct {
	Name                string    `json:"name"`
	DeviceID            string    `json:"device_id"`
	DeviceType          string    `json:"device_type"`
	Token               string    `json:"token"`
	TokenExpired        time.Time `json:"token_expired"`
	RefreshToken        string    `json:"refresh_token"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired"`
	IsLogin             bool      `json:"is_login"`
	UserLogin           string    `json:"user_login"`
}

// Validate ...
func (payload *AuthTokenPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}

	return
}

// ToPayloadAuthToken ...
func ToPayloadAuthToken(data sqlc.AuthToken) (response readAuthTokenPayload) {
	response = readAuthTokenPayload{
		Name:                data.Name,
		DeviceID:            data.DeviceID,
		DeviceType:          data.DeviceType,
		Token:               data.Token,
		TokenExpired:        data.TokenExpired,
		RefreshToken:        data.RefreshToken,
		RefreshTokenExpired: data.RefreshTokenExpired,
		IsLogin:             data.IsLogin,
	}

	if data.UserLogin.Valid {
		response.UserLogin = data.UserLogin.String
	}

	return
}
