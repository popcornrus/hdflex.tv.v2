package user_service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	request "go-boilerplate/internal/root/http/request/users"
	"go-boilerplate/internal/root/model"
	"go-boilerplate/internal/root/repository"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type UserService struct {
	log   *slog.Logger
	cache *cache.Cache

	userRepo repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	Update(context.Context, *model.User, request.UpdateRequest) error
	SignUp(context.Context, request.SignUpRequest) (*string, error)
	SignIn(context.Context, request.SignInRequest) (*string, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewUserService(
	log *slog.Logger,
	cache *cache.Cache,
	userRepo repository.UserRepositoryInterface,
) *UserService {
	return &UserService{
		log:      log,
		cache:    cache,
		userRepo: userRepo,
	}
}

func (s *UserService) Update(ctx context.Context, user *model.User, request request.UpdateRequest) error {
	const op = "UserService.Update"

	log := s.log.With(
		slog.String("op", op),
	)

	user.Email = request.Email
	user.Username = request.Username
	user.UpdatedAt = time.Now()

	err := s.userRepo.Update(ctx, user)
	if err != nil {
		log.Error("failed to update user", slog.Any("err", err))
		return err
	}

	s.log.Info("user was updated", slog.Any("user", user.Token))
	s.cache.Set(user.Token, *user, time.Hour*24)

	return nil
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	const op = "UserService.FindUserByEmail"

	log := s.log.With(
		slog.String("op", op),
	)

	entity, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		log.Error("failed to find user", err)
		return nil, err
	}

	return &entity, nil
}

func (s *UserService) SignUp(ctx context.Context, request request.SignUpRequest) (*string, error) {
	const op = "UserService.SignUp"

	log := s.log.With(
		slog.String("op", op),
	)

	userUUID := uuid.New().String()

	var entity *model.User

	crPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 16)
	if err != nil {
		log.Error("failed to encrypt password")
		return nil, errors.New("failed to encrypt password")
	}

	entity = &model.User{
		UUID:     userUUID,
		Username: request.Username,
		Email:    request.Email,
		Password: string(crPassword),
	}

	_, err = s.userRepo.Create(ctx, entity)
	if err != nil {
		log.Error("failed to create user", slog.Any("err", err))
		return nil, err
	}

	token := _hash(fmt.Sprintf("%s|%s", entity.UUID, time.Now()))
	entity.Token = token

	s.cache.Set(token, entity, time.Hour*24)

	return &token, nil
}

func (s *UserService) SignIn(ctx context.Context, request request.SignInRequest) (*string, error) {
	const op = "UserService.SignIn"

	log := s.log.With(
		slog.String("op", op),
	)

	entity, err := s.userRepo.FindUserByEmail(ctx, request.Email)
	if err != nil {
		log.Error("failed to find user", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(request.Password)); err != nil {
		log.Error("invalid password")
		return nil, err
	}

	token := _hash(fmt.Sprintf("%s|%s", entity.UUID, time.Now()))
	entity.Token = token

	s.cache.Set(token, entity, time.Hour*24)

	return &token, nil
}

func _hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
