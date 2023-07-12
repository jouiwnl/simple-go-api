// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: queries.sql

package sqlc

import (
	"context"
	"database/sql"
)

const CreateUser = `-- name: CreateUser :exec
insert into users(id, name, email, cpf, created_at) values ($1, $2, $3, $4, $5)
`

type CreateUserParams struct {
	ID        string
	Name      string
	Email     sql.NullString
	Cpf       sql.NullString
	CreatedAt sql.NullTime
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, CreateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Cpf,
		arg.CreatedAt,
	)
	return err
}

const ExistsUserByCpf = `-- name: ExistsUserByCpf :one
SELECT id, name, email, cpf, created_at, updated_at FROM users where cpf = $1 and id <> coalesce($2, 'any') limit 1
`

type ExistsUserByCpfParams struct {
	Cpf sql.NullString
	ID  string
}

func (q *Queries) ExistsUserByCpf(ctx context.Context, arg ExistsUserByCpfParams) (*User, error) {
	row := q.queryRow(ctx, q.existsUserByCpfStmt, ExistsUserByCpf, arg.Cpf, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Cpf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const ExistsUserByEmail = `-- name: ExistsUserByEmail :one
SELECT id, name, email, cpf, created_at, updated_at FROM users where email = $1 and id <> coalesce($2, 'any') limit 1
`

type ExistsUserByEmailParams struct {
	Email sql.NullString
	ID    string
}

func (q *Queries) ExistsUserByEmail(ctx context.Context, arg ExistsUserByEmailParams) (*User, error) {
	row := q.queryRow(ctx, q.existsUserByEmailStmt, ExistsUserByEmail, arg.Email, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Cpf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const GetPaginatedUsers = `-- name: GetPaginatedUsers :many
SELECT id, name, email, cpf, created_at, updated_at FROM users offset $1 limit $2
`

type GetPaginatedUsersParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) GetPaginatedUsers(ctx context.Context, arg GetPaginatedUsersParams) ([]*User, error) {
	rows, err := q.query(ctx, q.getPaginatedUsersStmt, GetPaginatedUsers, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Cpf,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetUserById = `-- name: GetUserById :one
SELECT id, name, email, cpf, created_at, updated_at FROM users where id = $1 limit 1
`

func (q *Queries) GetUserById(ctx context.Context, id string) (*User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, GetUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Cpf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const GetUsers = `-- name: GetUsers :many
SELECT id, name, email, cpf, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]*User, error) {
	rows, err := q.query(ctx, q.getUsersStmt, GetUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Cpf,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateUser = `-- name: UpdateUser :exec
update users set name = $1, email = $2, cpf = $3, updated_at = $4 where id = $5
`

type UpdateUserParams struct {
	Name      string
	Email     sql.NullString
	Cpf       sql.NullString
	UpdatedAt sql.NullTime
	ID        string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, UpdateUser,
		arg.Name,
		arg.Email,
		arg.Cpf,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
