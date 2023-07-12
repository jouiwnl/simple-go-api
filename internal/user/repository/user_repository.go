package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jouiwnl/simple-go-api/internal/application/commons"
	"github.com/jouiwnl/simple-go-api/internal/application/utils"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"time"
)

type UserRepository struct {
	db      *sqlc.DBTX
	queries *sqlc.Queries
	ctx     *context.Context
}

func NewUserRepository(db sqlc.DBTX, queries sqlc.Queries, ctx context.Context) *UserRepository {
	return &UserRepository{
		db:      &db,
		queries: &queries,
		ctx:     &ctx,
	}
}

func (r *UserRepository) GetUsers() []*sqlc.User {
	response, _ := r.queries.GetUsers(*r.ctx)

	return response
}

func (r *UserRepository) GetPaginatedUsers(pageable commons.Pageable) *commons.Page[dto.UserDto] {
	params := struct {
		Offset int32
		Limit  int32
	}{Offset: int32(pageable.Offset), Limit: int32(pageable.Limit)}

	var usersDto []dto.UserDto
	users, _ := r.queries.GetPaginatedUsers(*r.ctx, params)

	for _, user := range users {
		usersDto = append(usersDto, *dto.NewUserDtoByEntity(user))
	}

	response := commons.Page[dto.UserDto]{
		Total:    len(users),
		Content:  usersDto,
		Pageable: &pageable,
	}

	return &response
}

func (r *UserRepository) GetUserById(id string) (*sqlc.User, error) {
	result, err := r.queries.GetUserById(*r.ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) CreateUser(userDto *dto.UserDto) (string, error) {
	userId := uuid.New().String()
	createdAt := time.Now()

	user := sqlc.CreateUserParams{
		ID:   userId,
		Name: *userDto.Name,
		Email: sql.NullString{
			String: *userDto.Email,
			Valid:  true,
		},
		Cpf: sql.NullString{
			String: *userDto.Cpf,
			Valid:  true,
		},
		CreatedAt: sql.NullTime{
			Time:  createdAt,
			Valid: true,
		},
	}

	err := r.queries.CreateUser(*r.ctx, user)

	if err != nil {
		return "", err
	}

	return userId, nil
}

func (r *UserRepository) UpdateUser(userDto *dto.UserDto, id string) (string, error) {
	updatedAt := time.Now()

	user := sqlc.UpdateUserParams{
		ID:   id,
		Name: *userDto.Name,
		Email: sql.NullString{
			String: *userDto.Email,
			Valid:  true,
		},
		Cpf: sql.NullString{
			String: *userDto.Cpf,
			Valid:  true,
		},
		UpdatedAt: sql.NullTime{
			Time:  updatedAt,
			Valid: true,
		},
	}

	err := r.queries.UpdateUser(*r.ctx, user)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r *UserRepository) ExistsUserByCpf(id *string, cpf *string) (bool, error) {
	params := sqlc.ExistsUserByCpfParams{
		Cpf: sql.NullString{
			String: utils.NullSafeString(cpf),
			Valid:  true,
		},
		ID: utils.NullSafeString(id),
	}

	result, err := r.queries.ExistsUserByCpf(*r.ctx, params)

	if err != nil {
		return false, err
	}

	return result != nil, nil
}

func (r *UserRepository) ExistsUserByEmail(id *string, email *string) (bool, error) {
	params := sqlc.ExistsUserByEmailParams{
		Email: sql.NullString{
			String: utils.NullSafeString(email),
			Valid:  true,
		},
		ID: utils.NullSafeString(id),
	}

	result, err := r.queries.ExistsUserByEmail(*r.ctx, params)

	if err != nil {
		return false, err
	}

	return result != nil, nil
}
