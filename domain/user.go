package domain

import (
	"context"
	"time"
)

// User struct ...
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// UserRequestDTO ...
type UserRequestDTO struct {
	Name    string `json:"name" form:"name" query:"name"`
	Surname string `json:"surname" form:"surname" query:"surname"`
	Email   string `json:"email" form:"email" query:"email"`
}

// UserUsecases ...
type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

// UserRepository ...
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}
