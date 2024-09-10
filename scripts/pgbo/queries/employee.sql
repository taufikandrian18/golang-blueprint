-- name: InsertEmployee :one
INSERT INTO employee (guid, fullname, email, phone_number, date_of_birth, status, created_at, created_by)
VALUES (@guid, @fullname, @email, @phone_number, @date_of_birth, @status, (now() at time zone 'UTC'):: TIMESTAMP, @created_by)
RETURNING employee.*;

-- name: UpdateEmployee :one
UPDATE employee
SET fullname        = @fullname,
    email           = @email,
    phone_number    = @phone_number,
    date_of_birth   = @date_of_birth,
    updated_at      = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by      = @updated_by
WHERE guid = @guid
RETURNING employee.*;

-- name: GetEmployee :one
SELECT emp.guid,
       emp.employee_id,
       emp.fullname,
       emp.email,
       emp.phone_number,
       emp.date_of_birth,
       emp.status,
       emp.created_at,
       emp.created_by,
       emp.updated_at,
       emp.updated_by
FROM employee emp
WHERE emp.guid = @guid
  AND emp.status != 'deleted';

-- name: ListEmployee :many
SELECT emp.guid,
       emp.employee_id,
       emp.fullname,
       emp.email,
       emp.phone_number,
       emp.date_of_birth,
       emp.status,
       emp.created_at,
       emp.created_by,
       emp.updated_at,
       emp.updated_by
FROM employee emp
WHERE (CASE WHEN @set_guid::bool THEN emp.guid = @guid ELSE TRUE END)
  AND (CASE WHEN @set_fullname::bool THEN LOWER(emp.fullname) LIKE LOWER(@fullname) ELSE TRUE END)
  AND (CASE WHEN @set_email::bool THEN LOWER(emp.email) LIKE LOWER(@email) ELSE TRUE END)
  AND (CASE WHEN @set_date_of_birth::bool THEN TO_CHAR(emp.date_of_birth:: TIMESTAMP, 'DD/MM/YYYY') = TO_CHAR(@date_of_birth:: TIMESTAMP, 'DD/MM/YYYY') ELSE TRUE END)
  AND emp.status != 'deleted'
ORDER BY (CASE WHEN @order_param = 'id ASC' THEN emp.employee_id END) ASC,
         (CASE WHEN @order_param = 'id DESC' THEN emp.employee_id END) DESC,
         (CASE WHEN @order_param = 'guid ASC' THEN emp.guid END) ASC,
         (CASE WHEN @order_param = 'guid DESC' THEN emp.guid END) DESC,
         (CASE WHEN @order_param = 'fullname ASC' THEN emp.fullname END) ASC,
         (CASE WHEN @order_param = 'fullname DESC' THEN emp.fullname END) DESC,
         (CASE WHEN @order_param = 'email ASC' THEN emp.email END) ASC,
         (CASE WHEN @order_param = 'email DESC' THEN emp.email END) DESC,
         (CASE WHEN @order_param = 'phone_number ASC' THEN emp.phone_number END) ASC,
         (CASE WHEN @order_param = 'phone_number DESC' THEN emp.phone_number END) DESC,
         (CASE WHEN @order_param = 'date_of_birth ASC' THEN emp.date_of_birth END) ASC,
         (CASE WHEN @order_param = 'date_of_birth DESC' THEN emp.date_of_birth END) DESC,
         (CASE WHEN @order_param = 'created_at ASC' THEN emp.created_at END) ASC,
         (CASE WHEN @order_param = 'created_at DESC' THEN emp.created_at END) DESC,
         emp.created_at DESC
LIMIT @limit_data OFFSET @offset_page;

-- name: CountEmployee :one
SELECT COUNT(emp.employee_id)
FROM employee emp
WHERE (CASE WHEN @set_guid::bool THEN emp.guid = @guid ELSE TRUE END)
  AND (CASE WHEN @set_fullname::bool THEN LOWER(emp.fullname) LIKE LOWER(@fullname) ELSE TRUE END)
  AND (CASE WHEN @set_email::bool THEN LOWER(emp.email) LIKE LOWER(@email) ELSE TRUE END)
  AND (CASE WHEN @set_date_of_birth::bool THEN TO_CHAR(emp.date_of_birth:: TIMESTAMP, 'DD/MM/YYYY') = TO_CHAR(@date_of_birth:: TIMESTAMP, 'DD/MM/YYYY') ELSE TRUE END)
  AND emp.status != 'deleted';

-- name: UpdateEmployeeStatus :one
UPDATE employee
SET status     = @status,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE guid = @guid
  AND status != 'deleted'
RETURNING employee.*;