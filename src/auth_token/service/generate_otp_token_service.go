package service

import (
	"database/sql"

	"gitlab.com/wit-id/project-latihan/common/jwt"
)

type OTPInsertUserHandheldParams struct {
	Guid                   string         `json:"guid"`
	Name                   string         `json:"name"`
	ProfilePictureImageUrl sql.NullString `json:"profile_picture_image_url"`
	Phone                  sql.NullString `json:"phone"`
	Email                  string         `json:"email"`
	Gender                 string         `json:"gender"`
	Address                sql.NullString `json:"address"`
	Salt                   string         `json:"salt"`
	Password               string         `json:"password"`
	FcmToken               sql.NullString `json:"fcm_token"`
}

func (s *AuthTokenService) OtpToken(request OTPInsertUserHandheldParams, otp string) (jwtResponse jwt.ResponseJwtTokenOTP, err error) {
	jwtResponse, err = jwt.CreateJWTTokenOTPInsertUserHandheld(jwt.RequestJWTOTPInsertUserHandheldParams{
		Guid:                   request.Guid,
		Name:                   request.Name,
		ProfilePictureImageUrl: request.ProfilePictureImageUrl,
		Phone:                  request.Phone,
		Email:                  request.Email,
		Gender:                 request.Gender,
		Address:                request.Address,
		Salt:                   request.Salt,
		Password:               request.Password,
		FcmToken:               request.FcmToken,
	}, otp, s.cfg)
	if err != nil {
		return
	}

	return
}
