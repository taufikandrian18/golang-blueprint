// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: authentication.sql

package sqlc

import (
	"context"
	"database/sql"
)

const getAuthenticationByEmployeeID = `-- name: GetAuthenticationByEmployeeID :one
SELECT
    guid, id, employee_guid, username, password, forgot_password_token, forgot_password_expiry, is_active, last_login, status, created_at, created_by, updated_at, updated_by, salt
FROM
    authentication a
WHERE
    a.employee_guid = $1
    AND a.status != 'deleted'
`

func (q *Queries) GetAuthenticationByEmployeeID(ctx context.Context, employeeGuid sql.NullString) (Authentication, error) {
	row := q.db.QueryRowContext(ctx, getAuthenticationByEmployeeID, employeeGuid)
	var i Authentication
	err := row.Scan(
		&i.Guid,
		&i.ID,
		&i.EmployeeGuid,
		&i.Username,
		&i.Password,
		&i.ForgotPasswordToken,
		&i.ForgotPasswordExpiry,
		&i.IsActive,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Salt,
	)
	return i, err
}

const getAuthenticationByForgotPasswordToken = `-- name: GetAuthenticationByForgotPasswordToken :one
SELECT 
    a.guid,
    a.id,
    a.employee_guid,
    e.fullname AS employee_fullname,
    a.username,
    a.password,
    a.salt,
    a.status,
    a.forgot_password_token,
    a.forgot_password_expiry,
    a.created_at,
    a.created_by,
    a.updated_at,
    a.updated_by
FROM 
    authentication a
LEFT JOIN
    employee e ON a.employee_guid = e.guid
WHERE
    a.forgot_password_token = $1
    AND a.status != 'deleted'
LIMIT 1
`

type GetAuthenticationByForgotPasswordTokenRow struct {
	Guid                 string         `json:"guid"`
	ID                   sql.NullInt32  `json:"id"`
	EmployeeGuid         sql.NullString `json:"employee_guid"`
	EmployeeFullname     sql.NullString `json:"employee_fullname"`
	Username             string         `json:"username"`
	Password             string         `json:"password"`
	Salt                 sql.NullString `json:"salt"`
	Status               string         `json:"status"`
	ForgotPasswordToken  sql.NullString `json:"forgot_password_token"`
	ForgotPasswordExpiry sql.NullTime   `json:"forgot_password_expiry"`
	CreatedAt            sql.NullTime   `json:"created_at"`
	CreatedBy            string         `json:"created_by"`
	UpdatedAt            sql.NullTime   `json:"updated_at"`
	UpdatedBy            sql.NullString `json:"updated_by"`
}

func (q *Queries) GetAuthenticationByForgotPasswordToken(ctx context.Context, forgotPasswordToken sql.NullString) (GetAuthenticationByForgotPasswordTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthenticationByForgotPasswordToken, forgotPasswordToken)
	var i GetAuthenticationByForgotPasswordTokenRow
	err := row.Scan(
		&i.Guid,
		&i.ID,
		&i.EmployeeGuid,
		&i.EmployeeFullname,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.Status,
		&i.ForgotPasswordToken,
		&i.ForgotPasswordExpiry,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const getAuthenticationByID = `-- name: GetAuthenticationByID :one
SELECT 
    a.guid,
    a.id,
    a.employee_guid,
    e.fullname AS employee_fullname,
    a.username,
    a.password,
    a.salt,
    a.status,
    a.forgot_password_token,
    a.forgot_password_expiry,
    a.created_at,
    a.created_by,
    a.updated_at,
    a.updated_by
FROM 
    authentication a
LEFT JOIN
    employee e ON a.employee_guid = e.guid
WHERE
    a.guid = $1
    AND a.status != 'deleted'
LIMIT 1
`

type GetAuthenticationByIDRow struct {
	Guid                 string         `json:"guid"`
	ID                   sql.NullInt32  `json:"id"`
	EmployeeGuid         sql.NullString `json:"employee_guid"`
	EmployeeFullname     sql.NullString `json:"employee_fullname"`
	Username             string         `json:"username"`
	Password             string         `json:"password"`
	Salt                 sql.NullString `json:"salt"`
	Status               string         `json:"status"`
	ForgotPasswordToken  sql.NullString `json:"forgot_password_token"`
	ForgotPasswordExpiry sql.NullTime   `json:"forgot_password_expiry"`
	CreatedAt            sql.NullTime   `json:"created_at"`
	CreatedBy            string         `json:"created_by"`
	UpdatedAt            sql.NullTime   `json:"updated_at"`
	UpdatedBy            sql.NullString `json:"updated_by"`
}

func (q *Queries) GetAuthenticationByID(ctx context.Context, guid string) (GetAuthenticationByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthenticationByID, guid)
	var i GetAuthenticationByIDRow
	err := row.Scan(
		&i.Guid,
		&i.ID,
		&i.EmployeeGuid,
		&i.EmployeeFullname,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.Status,
		&i.ForgotPasswordToken,
		&i.ForgotPasswordExpiry,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const getAuthenticationByUsername = `-- name: GetAuthenticationByUsername :one
SELECT 
    a.guid,
    a.id,
    a.employee_guid,
    e.fullname AS employee_fullname,
    a.username,
    a.password,
    a.salt,
    a.status,
    a.forgot_password_token,
    a.forgot_password_expiry,
    a.created_at,
    a.created_by,
    a.updated_at,
    a.updated_by
FROM 
    authentication a
LEFT JOIN
    employee e ON a.employee_guid = e.guid
WHERE
    a.username = $1
    AND a.status != 'deleted'
LIMIT 1
`

type GetAuthenticationByUsernameRow struct {
	Guid                 string         `json:"guid"`
	ID                   sql.NullInt32  `json:"id"`
	EmployeeGuid         sql.NullString `json:"employee_guid"`
	EmployeeFullname     sql.NullString `json:"employee_fullname"`
	Username             string         `json:"username"`
	Password             string         `json:"password"`
	Salt                 sql.NullString `json:"salt"`
	Status               string         `json:"status"`
	ForgotPasswordToken  sql.NullString `json:"forgot_password_token"`
	ForgotPasswordExpiry sql.NullTime   `json:"forgot_password_expiry"`
	CreatedAt            sql.NullTime   `json:"created_at"`
	CreatedBy            string         `json:"created_by"`
	UpdatedAt            sql.NullTime   `json:"updated_at"`
	UpdatedBy            sql.NullString `json:"updated_by"`
}

func (q *Queries) GetAuthenticationByUsername(ctx context.Context, username string) (GetAuthenticationByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthenticationByUsername, username)
	var i GetAuthenticationByUsernameRow
	err := row.Scan(
		&i.Guid,
		&i.ID,
		&i.EmployeeGuid,
		&i.EmployeeFullname,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.Status,
		&i.ForgotPasswordToken,
		&i.ForgotPasswordExpiry,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const insertAuthentication = `-- name: InsertAuthentication :one
INSERT INTO authentication(
	guid, employee_guid, username, password, salt, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING authentication.guid, authentication.id, authentication.employee_guid, authentication.username, authentication.password, authentication.forgot_password_token, authentication.forgot_password_expiry, authentication.is_active, authentication.last_login, authentication.status, authentication.created_at, authentication.created_by, authentication.updated_at, authentication.updated_by, authentication.salt
`

type InsertAuthenticationParams struct {
	Guid         string         `json:"guid"`
	EmployeeGuid sql.NullString `json:"employee_guid"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	Salt         sql.NullString `json:"salt"`
	Status       string         `json:"status"`
	CreatedBy    string         `json:"created_by"`
}

func (q *Queries) InsertAuthentication(ctx context.Context, arg InsertAuthenticationParams) (Authentication, error) {
	row := q.db.QueryRowContext(ctx, insertAuthentication,
		arg.Guid,
		arg.EmployeeGuid,
		arg.Username,
		arg.Password,
		arg.Salt,
		arg.Status,
		arg.CreatedBy,
	)
	var i Authentication
	err := row.Scan(
		&i.Guid,
		&i.ID,
		&i.EmployeeGuid,
		&i.Username,
		&i.Password,
		&i.ForgotPasswordToken,
		&i.ForgotPasswordExpiry,
		&i.IsActive,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Salt,
	)
	return i, err
}

const recordAuthenticationLastLogin = `-- name: RecordAuthenticationLastLogin :exec
UPDATE authentication
SET
    last_login = (now() at time zone 'UTC')::TIMESTAMP
WHERE
    guid = $1
    AND status = 'active'
`

func (q *Queries) RecordAuthenticationLastLogin(ctx context.Context, guid string) error {
	_, err := q.db.ExecContext(ctx, recordAuthenticationLastLogin, guid)
	return err
}

const updateAuthenticationForgotPassword = `-- name: UpdateAuthenticationForgotPassword :exec
UPDATE authentication
SET
    forgot_password_token = $1,
    forgot_password_expiry = $2,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $3
WHERE
    guid = $4
`

type UpdateAuthenticationForgotPasswordParams struct {
	ForgotPasswordToken  sql.NullString `json:"forgot_password_token"`
	ForgotPasswordExpiry sql.NullTime   `json:"forgot_password_expiry"`
	UpdatedBy            sql.NullString `json:"updated_by"`
	Guid                 string         `json:"guid"`
}

func (q *Queries) UpdateAuthenticationForgotPassword(ctx context.Context, arg UpdateAuthenticationForgotPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthenticationForgotPassword,
		arg.ForgotPasswordToken,
		arg.ForgotPasswordExpiry,
		arg.UpdatedBy,
		arg.Guid,
	)
	return err
}

const updateAuthenticationPassword = `-- name: UpdateAuthenticationPassword :exec
UPDATE authentication
SET
    password = $1,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $2
WHERE
    guid = $3
`

type UpdateAuthenticationPasswordParams struct {
	Password  string         `json:"password"`
	UpdatedBy sql.NullString `json:"updated_by"`
	Guid      string         `json:"guid"`
}

func (q *Queries) UpdateAuthenticationPassword(ctx context.Context, arg UpdateAuthenticationPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthenticationPassword, arg.Password, arg.UpdatedBy, arg.Guid)
	return err
}

const updateAuthenticationUsername = `-- name: UpdateAuthenticationUsername :exec
UPDATE authentication
SET
    username = $1,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $2
WHERE
    guid = $3
`

type UpdateAuthenticationUsernameParams struct {
	Username  string         `json:"username"`
	UpdatedBy sql.NullString `json:"updated_by"`
	Guid      string         `json:"guid"`
}

func (q *Queries) UpdateAuthenticationUsername(ctx context.Context, arg UpdateAuthenticationUsernameParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthenticationUsername, arg.Username, arg.UpdatedBy, arg.Guid)
	return err
}

const updateAuthenticationUsernameByEmployeeID = `-- name: UpdateAuthenticationUsernameByEmployeeID :exec
UPDATE authentication
SET
    username = $1,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $2
WHERE
    employee_guid = $3
`

type UpdateAuthenticationUsernameByEmployeeIDParams struct {
	Username     string         `json:"username"`
	UpdatedBy    sql.NullString `json:"updated_by"`
	EmployeeGuid sql.NullString `json:"employee_guid"`
}

func (q *Queries) UpdateAuthenticationUsernameByEmployeeID(ctx context.Context, arg UpdateAuthenticationUsernameByEmployeeIDParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthenticationUsernameByEmployeeID, arg.Username, arg.UpdatedBy, arg.EmployeeGuid)
	return err
}
