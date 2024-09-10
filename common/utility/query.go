package utility

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/project-latihan/common/httpservice"
	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/project-latihan/toolkit/log"
)

func ReplaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	stmt = fmt.Sprintf(stmt, strings.Repeat(pattern, len))
	n := 0

	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}

	return strings.TrimSuffix(stmt, ",")
}

func Transaction(ctx context.Context, db *sql.DB, txFunc func(query *sqlc.Queries) (data interface{}, err error)) (data interface{}, err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin transaction")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	data, err = txFunc(sqlc.New(db).WithTx(tx))
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.FromCtx(ctx).Error(err, "failed rollback transaction")
			err = errors.WithStack(httpservice.ErrUnknownSource)

			return
		}

		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "failed commit transaction")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
