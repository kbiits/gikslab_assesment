// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package repoUser

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const findUserByUsername = `-- name: FindUserByUsername :one
SELECT id, name, email, username, password, profile, skill_id, created_at, updated_at FROM users u WHERE u.username = $1
`

func (q *Queries) FindUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Profile,
		&i.SkillID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUsersByMultipleIds = `-- name: FindUsersByMultipleIds :many
SELECT id, name, email, username, password, profile, skill_id, created_at, updated_at FROM users u WHERE id = ANY($1::int8[])
`

func (q *Queries) FindUsersByMultipleIds(ctx context.Context, dollar_1 []int64) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findUsersByMultipleIds, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.Profile,
			&i.SkillID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUsersByUsernames = `-- name: FindUsersByUsernames :many
SELECT id, name, email, username, password, profile, skill_id, created_at, updated_at FROM users u WHERE u.username = ANY($1::varchar[])
`

func (q *Queries) FindUsersByUsernames(ctx context.Context, dollar_1 []string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findUsersByUsernames, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.Profile,
			&i.SkillID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const registerUser = `-- name: RegisterUser :one
INSERT INTO
    users (name, email, username, password, profile, skill_id)
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, username, password, profile, skill_id, created_at, updated_at
`

type RegisterUserParams struct {
	Name     string
	Email    string
	Username string
	Password string
	Profile  string
	SkillID  sql.NullInt64
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, registerUser,
		arg.Name,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.Profile,
		arg.SkillID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Profile,
		&i.SkillID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}