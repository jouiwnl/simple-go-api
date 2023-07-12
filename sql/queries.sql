-- name: GetUsers :many
SELECT * FROM users;

-- name: GetPaginatedUsers :many
SELECT * FROM users offset $1 limit $2;

-- name: GetUserById :one
SELECT * FROM users where id = $1 limit 1;

-- name: ExistsUserByCpf :one
SELECT * FROM users where cpf = $1 and id <> coalesce($2, 'any') limit 1;

-- name: ExistsUserByEmail :one
SELECT * FROM users where email = $1 and id <> coalesce($2, 'any') limit 1;

-- name: CreateUser :exec
insert into users(id, name, email, cpf, created_at) values ($1, $2, $3, $4, $5);

-- name: UpdateUser :exec
update users set name = $1, email = $2, cpf = $3, updated_at = $4 where id = $5;
