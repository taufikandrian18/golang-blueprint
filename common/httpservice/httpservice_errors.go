package httpservice

import (
	"errors"

	"gitlab.com/wit-id/project-latihan/common/constants"
)

// error message.
var (
	ErrInternalServerError = errors.New("internal server error")

	ErrBadRequest               = errors.New("bad request payload")
	ErrInvalidAppKey            = errors.New("invalid app key")
	ErrUnknownSource            = errors.New("unknown error")
	ErrUnknownSourceWithMessage = errors.New("unknown error with message")
	ErrBadRequestWithMessage    = errors.New("error with message")

	ErrMissingHeaderData = errors.New("missing header data")

	ErrInvalidToken            = errors.New("invalid token")
	ErrUnauthorizedTokenData   = errors.New("unauthorized token data")
	ErrInvalidOTP              = errors.New("invalid otp")
	ErrInvalidOTPToken         = errors.New("invalid otp token")
	ErrInvalidPhoneNumberOTP   = errors.New("invalid phone number for this otp")
	ErrPasswordNotMatch        = errors.New("password not match")
	ErrConfirmPasswordNotMatch = errors.New("confirm password not match")
	ErrNoResultData            = errors.New("no result data")
	ErrNoResultDataWithMessage = errors.New("no result data with message")

	ErrUserAlreadyRegistered = errors.New("user is already registered")
	ErrUserNotFound          = errors.New("user not found")
	ErrUnauthorizedUser      = errors.New("unauthorized user")
	ErrInActiveUser          = errors.New("user not active")
	SuccessChangePassword    = errors.New("successfully changed password")

	ErrRoleNotFound = errors.New("role not found")
	ErrDataNotFound = errors.New("data not found")

	ErrInvalidromotionCode         = errors.New("invalid promotion code")
	ErrInsufficientQuantityVoucher = errors.New("insufficient quantities of voucher")
	ErrVoucherIsNotActive          = errors.New("voucher is not active")
	ErrVoucherIsExpired            = errors.New("voucher is expired")

	ErrNoDoctor    = errors.New("doctor is not found")
	ErrNoSchedules = errors.New("schedule is not found")

	ErrUserAlreadyCheckIn       = errors.New("this user is checked in already")
	ErrAttemptIsReachedMax      = errors.New("attempt is reached maximum")
	ErrAttemptIsAlreadyRecorded = errors.New("attempt is taken already")
	ErrAssesmentCantUpdate      = errors.New("assesment cant update while already taken")

	ErrInvalidPaymentID = errors.New("invalid payment id")

	ErrNIKAlreadyExist         = errors.New("nik already exist")
	ErrIDCardAlreadyExist      = errors.New("id card already exist")
	ErrNPWPAlreadyExist        = errors.New("npwp already exist")
	ErrEmailAlreadyExist       = errors.New("email already exist")
	ErrPhoneNumberAlreadyExist = errors.New("phone number already exist")
)

// error message.
var (
	MsgHeaderTokenNotFound = constants.ErrorResponse{
		ID: "Header `token` tidak ditemukan",
		EN: "Header `token` not found",
	}

	MsgHeaderRefreshTokenNotFound = constants.ErrorResponse{
		ID: "Header `refresh-token` tidak ditemukan",
		EN: "Header `refresh-token` not found",
	}

	MsgHeaderTokenUnauthorized = constants.ErrorResponse{
		ID: "Token tidak sah",
		EN: "Unauthorized token",
	}

	MsgHeaderRefreshTokenUnauthorized = constants.ErrorResponse{
		ID: "Refresh token tidak sah",
		EN: "Unauthorized refresh token",
	}

	MsgIsNotLogin = constants.ErrorResponse{
		ID: "Silahkan masuk terlebih dahulu",
		EN: "Please login first",
	}

	MsgUnauthorizedUser = constants.ErrorResponse{
		ID: "Pengguna tidak sah",
		EN: "Unauthorized user",
	}

	MsgUserNotActive = constants.ErrorResponse{
		ID: "User tidak aktif",
		EN: "User not active",
	}

	MsgInvalidIDParam = constants.ErrorResponse{
		ID: "parameter id tidak valid",
		EN: "invalid id parameter",
	}

	MsgInvalidCheckIN = constants.ErrorResponse{
		ID: "user sudah check in hari ini",
		EN: "this user is checked in already",
	}

	MsgAttemptReachedMax = constants.ErrorResponse{
		ID: "user untuk attempt module ini sudah melampaui batas",
		EN: "user for this module Attempt Reached Max",
	}

	MsgAttemptIsTaken = constants.ErrorResponse{
		ID: "user untuk attempt module ini sudah tersimpan di database",
		EN: "user for this module Attempt is taken already",
	}

	MsgCantEditAssesmentWhichAlreadyTaken = constants.ErrorResponse{
		ID: "user tidak dapat mengubah assesment ketika assesment sudah di gunakan",
		EN: "user cant update this assesment while this assesment already taken",
	}
)
