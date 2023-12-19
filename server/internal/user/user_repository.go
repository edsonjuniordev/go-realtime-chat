package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertedId int
	query := "INSERT INTO users (email, password, username) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password, user.Username).Scan(&lastInsertedId)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertedId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	query := "SELECT id, email, password, username FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Username)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}
