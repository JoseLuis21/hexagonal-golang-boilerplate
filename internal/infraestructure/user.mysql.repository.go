package infraestructure

import (
	"context"
	"database/sql"
	"fmt"
	"hexagonal-go/internal/domain"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

const (
	sqlUserTable = "users"
)

type sqlUser struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the domain.UserRepository interface.
func (r *UserRepository) Save(ctx context.Context, user domain.UserModel) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		Password:  user.Password(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}
