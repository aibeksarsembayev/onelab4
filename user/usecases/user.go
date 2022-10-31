package usecase

import (
	"context"
	"time"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

// Create user ...
func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// Get user by ID ...
func (u *userUsecase) GetByID(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Get All users ...
func (u *userUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update user ...
func (u *userUsecase) Update(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	err := u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete user ...
func (u *userUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
