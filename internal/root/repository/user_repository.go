package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-boilerplate/external/db"
	"go-boilerplate/internal/root/model"
	"time"
)

type UserRepository struct {
	db      db.MysqlInterface
	mongodb db.MongoDBInterface
}

type UserRepositoryInterface interface {
	Update(context.Context, *model.User) error
	Create(context.Context, *model.User) (*model.User, error)
	FindUserByUUID(context.Context, string) (model.User, error)
	FindUserByID(context.Context, int) (model.User, error)
	FindUserByEmail(context.Context, string) (model.User, error)
}

func NewUserRepository(
	db db.MysqlInterface,
	mongoDB db.MongoDBInterface,
) *UserRepository {
	return &UserRepository{
		db:      db,
		mongodb: mongoDB,
	}
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	const query = "UPDATE `users` SET `username` = ?, `email` = ?, `updated_at` = ? WHERE `id` = ?"

	user.UpdatedAt = time.Now()

	_, err := r.db.GetExecer().ExecContext(ctx, query, user.Username, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	const query = "INSERT INTO `users` (`uuid`, `username`, `email`, `password`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?)"

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := r.db.GetExecer().ExecContext(ctx, query, user.UUID, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (r *UserRepository) FindUserByUUID(ctx context.Context, uuid string) (model.User, error) {
	const query = "SELECT `id`, `uuid`, `username`, `password`, `email`, `updated_at` FROM `users` WHERE `uuid` = ?"

	row := r.db.GetExecer().QueryRowContext(ctx, query, uuid)

	user := model.User{}

	err := row.Scan(&user.ID, &user.UUID, &user.Username, &user.Password, &user.Email, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}

		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindUserByID(ctx context.Context, userID int) (model.User, error) {
	const op = "repository.user.GetUserByID"

	const query = "SELECT `id`, `uuid`, `username`, `password`, `email`, `updated_at` FROM `users` WHERE `id` = ?"

	row := r.db.GetExecer().QueryRowContext(ctx, query, userID)

	user := model.User{}

	err := row.Scan(&user.ID, &user.UUID, &user.Username, &user.Password, &user.Email, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("%s: %w", op, err)
		}

		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	const op = "repository.user.FindUserByEmail"

	const query = "SELECT `id`, `uuid`, `username`, `password`, `email`, `updated_at` FROM `users` WHERE `email` = ?"

	row := r.db.GetExecer().QueryRowContext(ctx, query, email)

	user := model.User{}

	err := row.Scan(&user.ID, &user.UUID, &user.Username, &user.Password, &user.Email, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("%s: %w", op, err)
		}

		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
