package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type SignUpReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

type SignUpRes struct {
	ID       int64  `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
}

type SignInReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type SignInRes struct {
	accessToken string
	ID          int64  `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
}

type Service interface {
	SignUp(c context.Context, req *SignUpReq) (*SignUpRes, error)
	SignIn(c context.Context, req *SignInReq) (*SignInRes, error)
}
