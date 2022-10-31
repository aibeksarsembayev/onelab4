package db

import (
	"context"
	"errors"
	"time"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
)

type dbUserRepository struct {
	// Conn
}

// NewDBUserRepository ...
func NewDBUserRepository() domain.UserRepository {
	return &dbUserRepository{}
}

// Create user in db ...
func (db *dbUserRepository) Create(ctx context.Context, user *domain.User) error {
	// Insert in db
	return nil
}

// Get user by ID from db ...
func (db *dbUserRepository) GetByID(ctx context.Context, id int) (domain.User, error) {
	// Select from db by id
	if id == 1 {
		user := domain.User{
			ID:      id,
			Name:    "Leela",
			Surname: "Turanga",
			Email:   "one@eye.com",
			Status:  false,
		}
		return user, nil
	} else {
		err := errors.New("no record")
		return domain.User{}, err
	}
}

// Get all users from db ..
func (db *dbUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{
		{
			ID:        0,
			Name:      "bender",
			Surname:   "unknown",
			Email:     "r3000@m.com",
			Status:    true,
			CreatedAt: time.Now().Add(1000 * time.Hour),
		},
		{
			ID:        1,
			Name:      "Leela",
			Surname:   "Turanga",
			Email:     "one@eye.com",
			Status:    false,
			CreatedAt: time.Now().Add(9999 * time.Hour)},
	}
	return users, nil
}

// Update user in db
func (db *dbUserRepository) Update(ctx context.Context, user *domain.User) error {
	if user.ID == 1 {
		return nil
	} else {
		err := errors.New("no record")
		return err
	}
}

// Delete user by id in db
func (db *dbUserRepository) Delete(ctx context.Context, id int) error {
	if id == 1 {
		return nil
	} else {
		err := errors.New("no record")
		return err
	}
}
