package jwt

import (
	"context"
	"database/sql"
	"time"

	"gitlab.com/wit-id/project-latihan/toolkit/log"

	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/toolkit/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type RequestJWTToken struct {
	AppName    string
	DeviceID   string
	DeviceType string
	IPAddress  string
}

type RequestJWTOTPInsertUserHandheldParams struct {
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

type RequestJWTOTPForgotPasswordUserHandheldParams struct {
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

type RequestJWTResetPasswordUserHandheldParams struct {
	Guid string `json:"guid"`
}

type ResponseJwtToken struct {
	AppName             string
	DeviceID            string
	DeviceType          string
	IPAddress           string
	Token               string
	TokenExpired        time.Time
	RefreshToken        string
	RefreshTokenExpired time.Time
}

type ResponseJwtTokenOTP struct {
	Token        string    `json:"token,omitempty"`
	TokenExpired time.Time `json:"token_expired"`
}

type ResponseJWTTokenResetPassword struct {
	Token        string    `json:"reset_password_token,omitempty"`
	TokenExpired time.Time `json:"reset_password_token_expired"`
}

// JWT token ...
func CreateJWTToken(ctx context.Context, cfg config.KVStore, request RequestJWTToken) (response ResponseJwtToken, err error) {
	tokenJwt := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	expiredToken := time.Now().Add(cfg.GetDuration("jwt.expired"))
	// The backend can also decode the token and get admin etc.
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["app_name"] = request.AppName
	claims["device_id"] = request.DeviceID
	claims["device_type"] = request.DeviceType
	claims["ip_address"] = request.IPAddress
	claims["exp"] = expiredToken.Unix()

	// The signing string should be secret (a generated UUID works too)
	token, err := tokenJwt.SignedString([]byte(cfg.GetString("jwt.key")))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed generate jwt token")
		err = errors.WithStack(httpservice.ErrInternalServerError)

		return
	}

	refreshTokenJwt := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	expiredRefreshToken := time.Now().Add(cfg.GetDuration("jwt.refresh_expired"))
	// The backend can also decode the token and get admin etc.
	rtClaims := refreshTokenJwt.Claims.(jwt.MapClaims)
	rtClaims["app_name"] = request.AppName
	rtClaims["device_id"] = request.DeviceID
	rtClaims["device_type"] = request.DeviceType
	rtClaims["ip_address"] = request.IPAddress
	rtClaims["exp"] = expiredRefreshToken.Unix()

	// The signing string should be secret (a generated UUID works too)
	refreshToken, err := refreshTokenJwt.SignedString([]byte(cfg.GetString("jwt.key")))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed generate refresh token")
		err = errors.WithStack(httpservice.ErrInternalServerError)

		return
	}

	response = ResponseJwtToken{
		AppName:             request.AppName,
		DeviceID:            request.DeviceID,
		DeviceType:          request.DeviceType,
		IPAddress:           request.IPAddress,
		Token:               token,
		TokenExpired:        expiredToken,
		RefreshToken:        refreshToken,
		RefreshTokenExpired: expiredRefreshToken,
	}

	return
}

func ClaimsJwtToken(ctx context.Context, cfg config.KVStore, token string) (response RequestJWTToken, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			log.FromCtx(ctx).Error(err, "Unexpected signing method: %v", token.Header["alg"])

			return nil, errors.WithStack(httpservice.ErrInvalidToken)
		}
		return []byte(cfg.GetString("jwt.key")), nil
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed parse token")
		err = errors.WithStack(httpservice.ErrInvalidToken)

		return
	}

	if jwtToken == nil {
		err = errors.WithStack(httpservice.ErrInvalidToken)
		return
	}

	if !jwtToken.Valid {
		err = errors.WithStack(httpservice.ErrInvalidToken)
		return
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		response = RequestJWTToken{
			AppName:    claims["app_name"].(string),
			DeviceID:   claims["device_id"].(string),
			DeviceType: claims["device_type"].(string),
			IPAddress:  claims["ip_address"].(string),
		}
	}

	return
}

// JWT for OTP...
func CreateJWTTokenOTPInsertUserHandheld(request RequestJWTOTPInsertUserHandheldParams, otp string, cfg config.KVStore) (response ResponseJwtTokenOTP, err error) {
	tokenJwt := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	expiredToken := time.Now().Add(time.Duration(cfg.GetInt64("jwt.expired-otp")))
	// The backend can also decode the token and get admin etc.
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["Guid"] = request.Guid
	claims["Name"] = request.Name
	claims["ProfilePictureImageUrl"] = request.ProfilePictureImageUrl.String
	claims["Phone"] = request.Phone.String
	claims["Email"] = request.Email
	claims["Gender"] = request.Gender
	claims["Address"] = request.Address.String
	claims["Salt"] = request.Salt
	claims["Password"] = request.Password
	claims["FcmToken"] = request.FcmToken.String
	claims["OTP"] = otp

	// The signing string should be secret (a generated UUID works too)
	token, err := tokenJwt.SignedString([]byte(cfg.GetString("jwt.key-otp")))
	if err != nil {
		err = errors.Wrap(err, "failed generate jwt token otp")
		return
	}

	response = ResponseJwtTokenOTP{
		Token:        token,
		TokenExpired: expiredToken,
	}

	return
}

func ClaimsJWTTokenOtpInsertUserHandheld(cfg config.KVStore, token string) (response RequestJWTOTPInsertUserHandheldParams, otp string, err error) {
	tokenOtp, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, errors.Wrapf(err, "Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.GetString("jwt.key-otp")), nil
	})
	if err != nil {
		return
	}

	if tokenOtp == nil {
		err = errors.WithStack(httpservice.ErrInvalidOTPToken)
		return
	}

	if !tokenOtp.Valid {
		err = errors.WithStack(httpservice.ErrInvalidOTPToken)
		return
	}

	if claims, ok := tokenOtp.Claims.(jwt.MapClaims); ok && tokenOtp.Valid {
		response = RequestJWTOTPInsertUserHandheldParams{
			Guid:     claims["Guid"].(string),
			Name:     claims["Name"].(string),
			Email:    claims["Email"].(string),
			Gender:   claims["Gender"].(string),
			Salt:     claims["Salt"].(string),
			Password: claims["Password"].(string),
		}

		if claims["ProfilePictureImageUrl"].(string) != "" {
			response.ProfilePictureImageUrl = sql.NullString{
				String: claims["ProfilePictureImageUrl"].(string),
				Valid:  true,
			}
		}

		if claims["Phone"].(string) != "" {
			response.Phone = sql.NullString{
				String: claims["Phone"].(string),
				Valid:  true,
			}
		}

		if claims["Address"].(string) != "" {
			response.Address = sql.NullString{
				String: claims["Address"].(string),
				Valid:  true,
			}
		}

		if claims["FcmToken"].(string) != "" {
			response.FcmToken = sql.NullString{
				String: claims["FcmToken"].(string),
				Valid:  true,
			}
		}

		otp = claims["OTP"].(string)
	}

	return
}

func CreateJWTTokenOTPForgotPasswordUserHandheld(cfg config.KVStore, request RequestJWTOTPForgotPasswordUserHandheldParams, otp string) (response ResponseJwtTokenOTP, err error) {
	tokenJwt := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	expiredToken := time.Now().Add(time.Duration(cfg.GetInt64("jwt.expired-otp")))
	// The backend can also decode the token and get admin etc.
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["Guid"] = request.Guid
	claims["Name"] = request.Name
	claims["ProfilePictureImageUrl"] = request.ProfilePictureImageUrl.String
	claims["Phone"] = request.Phone.String
	claims["Email"] = request.Email
	claims["Gender"] = request.Gender
	claims["Address"] = request.Address.String
	claims["Salt"] = request.Salt
	claims["Password"] = request.Password
	claims["FcmToken"] = request.FcmToken.String
	claims["OTP"] = otp

	// The signing string should be secret (a generated UUID works too)
	token, err := tokenJwt.SignedString([]byte(cfg.GetString("jwt.key-otp")))
	if err != nil {
		err = errors.Wrap(err, "failed generate jwt token otp")
		return
	}

	response = ResponseJwtTokenOTP{
		Token:        token,
		TokenExpired: expiredToken,
	}

	return
}

func ClaimsJWTTokenOtpForgotPasswordUserHandheld(cfg config.KVStore, token string) (response RequestJWTOTPForgotPasswordUserHandheldParams, otp string, err error) {
	tokenOtp, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, errors.Wrapf(err, "Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.GetString("jwt.key-otp")), nil
	})
	if err != nil {
		return
	}

	if tokenOtp == nil {
		err = errors.WithStack(httpservice.ErrInvalidOTPToken)
		return
	}

	if !tokenOtp.Valid {
		err = errors.WithStack(httpservice.ErrInvalidOTPToken)
		return
	}

	if claims, ok := tokenOtp.Claims.(jwt.MapClaims); ok && tokenOtp.Valid {
		response = RequestJWTOTPForgotPasswordUserHandheldParams{
			Guid:     claims["Guid"].(string),
			Name:     claims["Name"].(string),
			Email:    claims["Email"].(string),
			Gender:   claims["Gender"].(string),
			Salt:     claims["Salt"].(string),
			Password: claims["Password"].(string),
		}

		if claims["ProfilePictureImageUrl"].(string) != "" {
			response.ProfilePictureImageUrl = sql.NullString{
				String: claims["ProfilePictureImageUrl"].(string),
				Valid:  true,
			}
		}

		if claims["Phone"].(string) != "" {
			response.Phone = sql.NullString{
				String: claims["Phone"].(string),
				Valid:  true,
			}
		}

		if claims["Address"].(string) != "" {
			response.Address = sql.NullString{
				String: claims["Address"].(string),
				Valid:  true,
			}
		}

		if claims["FcmToken"].(string) != "" {
			response.FcmToken = sql.NullString{
				String: claims["FcmToken"].(string),
				Valid:  true,
			}
		}

		otp = claims["OTP"].(string)
	}

	return
}

// JWT for Reset Password...
func CreateJWTTokenResetPasswordUserHandheld(cfg config.KVStore, request RequestJWTResetPasswordUserHandheldParams) (response ResponseJWTTokenResetPassword, err error) {
	tokenJwt := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	expiredToken := time.Now().Add(time.Duration(cfg.GetInt64("jwt.expired-otp")))
	// The backend can also decode the token and get admin etc.
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["Guid"] = request.Guid

	// The signing string should be secret (a generated UUID works too)
	token, err := tokenJwt.SignedString([]byte(cfg.GetString("jwt.key-otp")))
	if err != nil {
		err = errors.Wrap(err, "failed generate jwt token otp")
		return
	}

	response = ResponseJWTTokenResetPassword{
		Token:        token,
		TokenExpired: expiredToken,
	}

	return
}

func ClaimsJWTTokenResetPasswordUserHandheld(cfg config.KVStore, token string) (response RequestJWTResetPasswordUserHandheldParams, err error) {
	tokenResetPassword, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, errors.Wrapf(err, "Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.GetString("jwt.key-otp")), nil
	})
	if err != nil {
		return
	}

	if tokenResetPassword == nil {
		err = errors.WithStack(httpservice.ErrInvalidOTPToken)
		return
	}

	if !tokenResetPassword.Valid {
		err = errors.WithStack(httpservice.ErrInvalidOTPToken)
		return
	}

	if claims, ok := tokenResetPassword.Claims.(jwt.MapClaims); ok && tokenResetPassword.Valid {
		response = RequestJWTResetPasswordUserHandheldParams{
			Guid: claims["Guid"].(string),
		}
	}

	return
}
