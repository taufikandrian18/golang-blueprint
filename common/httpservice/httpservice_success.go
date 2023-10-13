package httpservice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/project-latihan/common/constants"
)

var Message = constants.ErrorResponse{
	ID: "Sukses",
	EN: "Success",
}

type Response struct {
	Data        interface{}             `json:"data"`
	CurrentPage int                     `json:"current_page,omitempty"`
	Limit       int                     `json:"limit,omitempty"`
	TotalPage   int                     `json:"total_page,omitempty"`
	TotalData   int64                   `json:"total_data,omitempty"`
	Message     constants.ErrorResponse `json:"message"`
}

//	if err != nil {
//		Message = constants.ErrorResponse{
//			ID: err.Error(),
//			EN: err.Error(),
//		}
//	}
func ResponseData(ctx echo.Context, data interface{}, err error) error {
	return ctx.JSONPretty(http.StatusOK, Response{
		Data:    data,
		Message: Message,
	}, "")
}

func ResponsePagination(ctx echo.Context, data interface{}, err error, page int, limit int, totaPage int, totalData int) error {
	if err != nil {
		Message = constants.ErrorResponse{
			ID: err.Error(),
			EN: err.Error(),
		}
	}

	return ctx.JSONPretty(http.StatusOK, Response{
		Data:        data,
		CurrentPage: page,
		Limit:       limit,
		TotalPage:   totaPage,
		TotalData:   int64(totalData),
		Message:     Message,
	}, "")
}
