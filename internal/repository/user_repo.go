package db

import (
	"context"
	"strings"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type dbUserRepository struct {
	dbpool *sqlx.DB
}

// NewDBUserRepository ...
func NewDBUserRepository(dbpool *sqlx.DB) domain.UserRepository {
	return &dbUserRepository{
		dbpool: dbpool,
	}
}

// Create user in db ...
func (db *dbUserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := db.dbpool.NamedExec(`INSERT INTO "user" (name, surname, email, status, created_at) VALUES (:name, :surname, :email, :status, :created_at)`, &user)
	if err != nil {
		return err
	}
	return nil
}

// Get user by ID from db ...
func (db *dbUserRepository) GetByID(ctx context.Context, id int) (domain.User, error) {
	user := domain.User{}
	err := db.dbpool.Get(&user, `SELECT * FROM "user" WHERE "id"=$1`, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Get all users from db ..
func (db *dbUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}
	err := db.dbpool.Select(&users, `SELECT * FROM "user" ORDER by id ASC`)

	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

// Update user in db
func (db *dbUserRepository) Update(ctx context.Context, user *domain.User) error {
	_, err := db.GetByID(ctx, user.ID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return err
		}
	}

	_, err = db.dbpool.NamedExec(`UPDATE "user" SET (name, surname, email, status) = (:name, :surname, :email, :status) WHERE "id" = :id`, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete user by id in db
func (db *dbUserRepository) Delete(ctx context.Context, id int) error {
	_, err := db.GetByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return err
		}
	}

	_, err = db.dbpool.Exec(`DELETE FROM "user" WHERE "id" = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
