package echohttp

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/project-latihan/common/constants"
)

func handleLanguage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Request().Header.Get("Accept-Language") != constants.LanguageEN && ctx.Request().Header.Get("Accept-Language") != constants.LanguageID {
			ctx.Request().Header.Set("Accept-Language", constants.LanguageEN)
		}

		return next(ctx)
	}
}
