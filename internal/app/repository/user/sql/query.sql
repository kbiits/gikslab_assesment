-- name: RegisterUser :one
INSERT INTO
    users (name, email, username, password, profile, skill_id)
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: FindUserByUsername :one
SELECT * FROM users u WHERE u.username = $1;

-- name: FindUsersByUsernames :many
SELECT * FROM users u WHERE u.username = ANY($1::varchar[]);

-- name: FindUsersByMultipleIds :many
SELECT * FROM users u WHERE id = ANY($1::int8[]);