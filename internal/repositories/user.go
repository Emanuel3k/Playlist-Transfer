package repositories

import (
	"database/sql"
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
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
	// toodo: implement this

	return nil, nil
}

func (r *UserRepository) Create(user *domain.User) *web.AppError {
	// toodo: implement this

	return nil
}
