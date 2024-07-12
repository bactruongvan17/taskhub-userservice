package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bactruongvan17/taskhub-userservice/src/pkg/model"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/repo"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/request"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/response"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo repo.PGInterface
}

func NewAuthService(repo repo.PGInterface) AuthServiceInterface {
	return &AuthService{repo: repo}
}

type AuthServiceInterface interface {
	SingUp(ctx context.Context, req request.SignUpRequest) error
	SignIn(ctx context.Context, req request.SignInRequest) (*response.SignInReponse, error)
}

func (s *AuthService) SingUp(ctx context.Context, req request.SignUpRequest) error {
	user, err := s.repo.GetUserByEmail(ctx, req.Email, nil)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("system error")
	}

	if user != nil {
		return errors.New("email has exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return errors.New("system error")
	}

	model := &model.User{
		Email:    req.Email,
		Password: utils.StringP(string(password)),
		FullName: req.FullName,
	}

	err = s.repo.CreateUser(ctx, model, nil)
	if err != nil {
		return errors.New("system error")
	}

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, req request.SignInRequest) (*response.SignInReponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email, nil)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("system error")
	}

	if user == nil {
		return nil, errors.New("account invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("account invalid")
	}

	token, err := s.createToken(*user)
	if err != nil {
		return nil, errors.New("system error: " + err.Error())
	}

	res := &response.SignInReponse{
		AccessToken: token,
		User: response.UserInfo{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
		},
	}

	return res, nil
}

func (s *AuthService) createToken(user model.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"iss": "taskhub-userservice",
		"aud": "user",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	log.Println(claims)

	return claims.SignedString([]byte("secret-key"))
}
