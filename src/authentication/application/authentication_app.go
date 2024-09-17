package application

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/constants"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/common/jwt"
	"gitlab.com/wit-id/project-latihan/src/authentication/service"
	"gitlab.com/wit-id/project-latihan/src/middleware"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/config"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func AddRouteAuthentication(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewAuthenticationService(s.GetDB(), cfg)
	mddw := middleware.NewEnsureToken(s.GetDB(), cfg)

	authApp := e.Group("/auth")
	authApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "auth app ok")
	})

	authApp.Use(mddw.ValidateToken)

	authApp.POST("/login", login(svc))
	authApp.POST("/logout", logout(svc))
	authApp.POST("/forgot_password_request", forgotPasswordRequest(svc))
	authApp.POST("/forgot_password_submit/:forgot_password_token", forgotPasswordSubmit(svc))
	authApp.POST("/profile", getProfile(svc), mddw.ValidateUserLogin)
	authApp.POST("/register", registerUser(svc))
	authApp.POST("/change_password", changeSelfPassword(svc), mddw.ValidateUserLogin)
}

func login(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.LoginPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		data, err := svc.Login(ctx.Request().Context(), request, ctx.Get("token-data").(jwt.RequestJWTToken))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadAuthentication(sqlc.GetAuthenticationByIDRow(data)), nil)
	}
}

func logout(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// var request payload.LogoutPayload
		// if err := ctx.Bind(&request); err != nil {
		// 	log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
		// 	return errors.WithStack(httpservice.ErrBadRequest)
		// }

		// if err := request.Validate(ctx.Request().Context()); err != nil {
		// 	return err
		// }

		err := svc.LogoutToken(ctx.Request().Context(), ctx.Get("token-data").(jwt.RequestJWTToken))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, nil, nil)
	}
}

func getProfile(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		userInfo := ctx.Get(constants.MddwUserBackoffice).(sqlc.GetAuthenticationByIDRow)

		profile, err := svc.GetProfileFromToken(ctx.Request().Context(), userInfo.EmployeeGuid.String)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadEmployee(profile), nil)
	}
}

func forgotPasswordRequest(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var request payload.ForgotPasswordRequestPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		err := svc.ForgotPasswordRequest(ctx.Request().Context(), request.Email)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, true, nil)
	}
}

func forgotPasswordSubmit(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		forgotPasswordToken := ctx.Param("forgot_password_token")
		if forgotPasswordToken == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.ForgotPasswordSubmitPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		if request.Password != request.PasswordConfirmation {
			return httpservice.ErrBadRequest
		}

		err := svc.ForgotPasswordChange(ctx.Request().Context(),
			forgotPasswordToken,
			request.Password)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, true, nil)
	}
}

func changeSelfPassword(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		userInfo := ctx.Get(constants.MddwUserBackoffice).(sqlc.GetAuthenticationByIDRow)
		profile, err := svc.GetProfileFromToken(ctx.Request().Context(), userInfo.EmployeeGuid.String)
		if err != nil {
			return err
		}

		var request payload.ChangePassword
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		err = svc.ChangeEmployeePassword(ctx.Request().Context(), profile.Email, request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, true, nil)
	}
}

func registerUser(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.RegisterEmployeePayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(c.Request().Context()); err != nil {
			return err
		}

		err := svc.RegisterEmployee(c.Request().Context(), request.ToEntity(), request)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, true, nil)
	}
}
