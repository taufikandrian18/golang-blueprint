package echohttp

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/constants"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/toolkit/config"
)

func handleEchoError(_ config.KVStore) echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		var echoError *echo.HTTPError

		// if *echo.HTTPError, let echokit middleware handles it
		if errors.As(err, &echoError) {
			return
		}

		statusCode := http.StatusInternalServerError
		// message := "mohon maaf, terjadi kesalahan pada server"
		var message interface{}

		message = ErrInternalServerErrror

		switch {
		case errors.Is(err, httpservice.ErrBadRequest):
			statusCode = http.StatusBadRequest
			message = ErrBadRequest

		case errors.Is(err, httpservice.ErrInvalidAppKey):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidAppKey

		case errors.Is(err, httpservice.ErrUnknownSource):
			statusCode = http.StatusUnprocessableEntity
			message = ErrUnknownSource

		case errors.Is(err, httpservice.ErrMissingHeaderData):
			statusCode = http.StatusBadRequest
			message = ErrMissingHeaderData

		case errors.Is(err, httpservice.ErrInvalidToken):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidToken

		case errors.Is(err, httpservice.ErrUserAlreadyCheckIn):
			statusCode = http.StatusBadRequest
			message = ErrUserAlreadyCheckIn

		case errors.Is(err, httpservice.ErrAttemptIsReachedMax):
			statusCode = http.StatusBadRequest
			message = ErrUserModuleAttemptReachedMax

		case errors.Is(err, httpservice.ErrAttemptIsAlreadyRecorded):
			statusCode = http.StatusBadRequest
			message = ErrUserModuleAttemptIsTaken

		case errors.Is(err, httpservice.ErrAssesmentCantUpdate):
			statusCode = http.StatusBadRequest
			message = ErrEditAssesmentWhichAlreadyTaken

		case errors.Is(err, httpservice.ErrUnauthorizedTokenData):
			statusCode = http.StatusUnauthorized
			message = ErrUnauthorizedTokenData

		case errors.Is(err, httpservice.ErrInvalidOTP):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidOTP

		case errors.Is(err, httpservice.ErrInvalidOTPToken):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidOTPToken

		case errors.Is(err, httpservice.ErrInvalidPhoneNumberOTP):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidPhoneNumberOTP

		case errors.Is(err, httpservice.ErrPasswordNotMatch):
			statusCode = http.StatusUnauthorized
			message = ErrPasswordNotMatch

		case errors.Is(err, httpservice.ErrConfirmPasswordNotMatch):
			statusCode = http.StatusBadRequest
			message = ErrPasswordNotMatch

		case errors.Is(err, httpservice.ErrNoResultData):
			statusCode = http.StatusNotFound
			message = ErrNoResultData

		case errors.Is(err, httpservice.ErrUserAlreadyRegistered):
			statusCode = http.StatusConflict
			message = ErrUserAlreadyRegistered

		case errors.Is(err, httpservice.ErrUserNotFound):
			statusCode = http.StatusUnauthorized
			message = ErrUserNotFound

		case errors.Is(err, httpservice.ErrUnauthorizedUser):
			statusCode = http.StatusUnauthorized
			message = ErrUnauthorizedUser

		case errors.Is(err, httpservice.ErrInActiveUser):
			statusCode = http.StatusUnauthorized
			message = ErrInactiveUser

		case errors.Is(err, httpservice.SuccessChangePassword):
			statusCode = http.StatusOK
			message = SuccessChangedPassword

		case errors.Is(err, httpservice.ErrRoleNotFound):
			statusCode = http.StatusConflict
			message = ErrRoleNotFound

		case errors.Is(err, httpservice.ErrInvalidromotionCode):
			statusCode = http.StatusForbidden
			message = ErrInvalidPromotionCode

		case errors.Is(err, httpservice.ErrInsufficientQuantityVoucher):
			statusCode = http.StatusConflict
			message = ErrInsufficientQuantityVoucher

		case errors.Is(err, httpservice.ErrVoucherIsNotActive):
			statusCode = http.StatusConflict
			message = ErrVoucherIsNotActive

		case errors.Is(err, httpservice.ErrVoucherIsExpired):
			statusCode = http.StatusConflict
			message = ErrVoucherIsExpired

		case errors.Is(err, httpservice.ErrInvalidPaymentID):
			statusCode = http.StatusConflict
			message = ErrInvalidPaymentID

		case errors.Is(err, httpservice.ErrBadRequestWithMessage):
			statusCode = http.StatusBadRequest

			langError := strings.Split(err.Error(), "|")
			message = constants.ErrorResponse{
				ID: strings.ReplaceAll(langError[1], ": error with message", ""),
				EN: langError[0],
			}
		case errors.Is(err, httpservice.ErrNoResultDataWithMessage):
			statusCode = http.StatusNotFound

			langError := strings.Split(err.Error(), "|")
			message = constants.ErrorResponse{
				ID: strings.ReplaceAll(langError[1], ": no result data with message", ""),
				EN: langError[0],
			}
		case errors.Is(err, httpservice.ErrNIKAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrNIKAlreadyExist
		case errors.Is(err, httpservice.ErrIDCardAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrIDCardAlreadyExist
		case errors.Is(err, httpservice.ErrNPWPAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrNPWPAlreadyExist
		case errors.Is(err, httpservice.ErrEmailAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrEmailAlreadyExist
		case errors.Is(err, httpservice.ErrPhoneNumberAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrPhoneNumberAlreadyExist
		case errors.Is(err, httpservice.ErrUnknownSourceWithMessage):
			statusCode = http.StatusUnprocessableEntity

			langError := strings.Split(err.Error(), "|")
			message = constants.ErrorResponse{
				ID: strings.ReplaceAll(langError[1], ": unknown error with message", ""),
				EN: langError[0],
			}
		}

		_ = ctx.JSON(statusCode, echo.NewHTTPError(statusCode, message))
	}
}
