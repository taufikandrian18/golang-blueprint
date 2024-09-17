-- name: InsertAuthentication :one
INSERT INTO authentication(
	guid, employee_guid, username, password, salt, status, created_by)
	VALUES (@guid, @employee_guid, @username, @password, @salt, @status, @created_by)
RETURNING authentication.*;

-- name: GetAuthenticationByUsername :one
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
    a.username = @username
    AND a.status != 'deleted'
LIMIT 1;

-- name: UpdateAuthenticationPassword :exec
UPDATE authentication
SET
    password = @password,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE
    guid = @guid;

-- name: UpdateAuthenticationUsername :exec
UPDATE authentication
SET
    username = @username,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE
    guid = @guid;

-- name: UpdateAuthenticationForgotPassword :exec
UPDATE authentication
SET
    forgot_password_token = @forgot_password_token,
    forgot_password_expiry = @forgot_password_expiry,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE
    guid = @guid;

-- name: GetAuthenticationByForgotPasswordToken :one
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
    a.forgot_password_token = @forgot_password_token
    AND a.status != 'deleted'
LIMIT 1;

-- name: GetAuthenticationByID :one
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
    a.guid = @guid
    AND a.status != 'deleted'
LIMIT 1;

-- name: GetAuthenticationByEmployeeID :one
SELECT
    *
FROM
    authentication a
WHERE
    a.employee_guid = @employee_guid
    AND a.status != 'deleted';

-- name: RecordAuthenticationLastLogin :exec
UPDATE authentication
SET
    last_login = (now() at time zone 'UTC')::TIMESTAMP
WHERE
    guid = @guid
    AND status = 'active';

-- name: UpdateAuthenticationUsernameByEmployeeID :exec
UPDATE authentication
SET
    username = @username,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE
    employee_guid = @employee_guid;