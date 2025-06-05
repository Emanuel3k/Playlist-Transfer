package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, *web.AppError) {
	query := `SELECT id, first_name, last_name, email, password FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, web.InternalServerError(fmt.Errorf("error scanning row: %w", err))
	}

	return &user, nil
}

func (r *UserRepository) Create(user *domain.User) *web.AppError {
	uid := uuid.NewString()
	user.ID = &uid

	query := `INSERT INTO users (id, first_name, last_name, email, password) values ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return web.InternalServerError(fmt.Errorf("error executing query: %w", err))
	}

	return nil
}
