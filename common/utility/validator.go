package utility

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func ValidateStruct(ctx context.Context, req interface{}) (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(req); err != nil {
		log.FromCtx(ctx).Error(err, "failed parsing payload")
		err = errors.WithStack(httpservice.ErrBadRequest)
		return
	}

	return
}
