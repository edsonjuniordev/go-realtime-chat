package user

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/edsonjuniordev/util"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) SignUp(c context.Context, req *SignUpReq) (*SignUpRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return &SignUpRes{}, err
	}

	user := &User{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}

	result, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return &SignUpRes{}, err
	}

	response := &SignUpRes{
		ID:       result.ID,
		Email:    result.Email,
		Username: result.Username,
	}

	return response, nil
}

type MyJWTClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) SignIn(c context.Context, req *SignInReq) (*SignInRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &SignInRes{}, err
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return &SignInRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return &SignInRes{}, err
	}

	return &SignInRes{
		accessToken: ss,
		Username:    user.Username,
		ID:          user.ID,
	}, nil
}
