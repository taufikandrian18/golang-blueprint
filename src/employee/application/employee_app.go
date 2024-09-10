package application

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/src/employee/service"
	"gitlab.com/wit-id/project-latihan/src/repository/payload"
	"gitlab.com/wit-id/project-latihan/toolkit/config"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func AddRouteEmployee(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewEmployeeService(s.GetDB(), cfg)

	employeeApp := e.Group("/employee")
	employeeApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "employee app ok")
	})

	employeeApp.POST("", createEmployee(svc))
	employeeApp.POST("/list", listEmployee(svc))
	employeeApp.GET("/detail/:guid", getEmployeeByGuid(svc))
	employeeApp.PUT("/:guid", updateEmployee(svc))
	employeeApp.DELETE("/:guid", deleteEmployee(svc))
}

func createEmployee(svc *service.EmployeeService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertEmployeePayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		banner, err := svc.CreateEmployee(ctx.Request().Context(), request.ToEntity())
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadEmployee(banner), nil)
	}
}

func listEmployee(svc *service.EmployeeService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.ListEmployeePayload

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Validate request
		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		listData, totalData, err := svc.ListEmployee(ctx.Request().Context(), request.ToEntity())
		if err != nil {
			return err
		}

		// TOTAL PAGE
		totalPage := math.Ceil(float64(totalData) / float64(request.Limit))

		return httpservice.ResponsePagination(ctx,
			payload.ToPayloadListEmployee(listData),
			nil, int(request.Offset),
			int(request.Limit),
			int(totalPage),
			int(totalData))
	}
}

func getEmployeeByGuid(svc *service.EmployeeService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		guid := ctx.Param("guid")
		if guid == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		banner, err := svc.GetEmployeeByGUID(ctx.Request().Context(), guid)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadEmployee(banner), nil)
	}
}

func updateEmployee(svc *service.EmployeeService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		guid := ctx.Param("guid")
		if guid == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		var request payload.UpdateEmployeePayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		banner, err := svc.UpdateEmployee(ctx.Request().Context(),
			request.ToEntity(guid))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadEmployee(banner), nil)
	}
}

func deleteEmployee(svc *service.EmployeeService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		guid := ctx.Param("guid")
		if guid == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		err := svc.DeleteEmployee(ctx.Request().Context(), guid)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, nil, nil)
	}
}
