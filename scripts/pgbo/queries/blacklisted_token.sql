-- name: InsertBlacklistedToken :one
INSERT INTO blacklisted_token(
	token, type)
	VALUES (@token, @type)
RETURNING blacklisted_token.*;

-- name: GetBlacklistedToken :one
SELECT 
    b.token,
    b.type,
    b.created_at
FROM 
    blacklisted_token b
WHERE
    b.token = @token
LIMIT 1;