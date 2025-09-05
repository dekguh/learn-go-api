package service

import (
	"errors"
	"os"
	"time"

	"github.com/dekguh/learn-go-api/internal/api/model"
	"github.com/dekguh/learn-go-api/internal/api/repository"
	"github.com/dekguh/learn-go-api/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(ID uint) (*model.User, error)
	RegisterUser(email, name, password string) (*model.User, error)
	LoginUser(email, password string, ctx *gin.Context) (*model.LoginUserResponse, error)
	RefreshToken(ctx *gin.Context) (*model.RefreshTokenResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func (service *userService) GetUserByEmail(email string) (*model.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	result, err := service.repo.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, errors.New("failed to get user by email")
	}

	return &model.User{
		ID:        result.ID,
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *userService) GetUserById(ID uint) (*model.User, error) {
	if ID == 0 {
		return nil, errors.New("id is required")
	}

	result, err := service.repo.FindById(ID)
	if err != nil {
		return nil, errors.New("failed to get user by id")
	}

	if result == nil {
		return nil, errors.New("user not found")
	}

	return &model.User{
		ID:        result.ID,
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *userService) RegisterUser(email, name, password string) (*model.User, error) {
	if exists, _ := service.repo.FindByEmail(email); exists != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	if err := service.repo.Create(user); err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *userService) LoginUser(email, password string, ctx *gin.Context) (*model.LoginUserResponse, error) {
	user, err := service.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwt.GenerateJwt(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate jwt")
	}
	refreshToken, err := jwt.GenerateRefreshJwt(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate refresh jwt")
	}

	ctx.SetCookie(
		os.Getenv("JWT_REFRESH_KEY"),
		refreshToken,
		int(1*24*time.Hour.Seconds()),
		"/",
		os.Getenv("COOKIE_HOST"),
		false,
		true,
	)
	return &model.LoginUserResponse{
		Token: token,
		User: &model.User{
			Email:     user.Email,
			Name:      user.Name,
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (service *userService) RefreshToken(ctx *gin.Context) (*model.RefreshTokenResponse, error) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		return nil, errors.New("refresh token is required")
	}

	claims, err := jwt.ParseJwt(refreshToken)
	if err != nil {
		return nil, errors.New("failed to parse refresh token")
	}

	newToken, err := jwt.GenerateJwt(claims.UserID, claims.Email)
	if err != nil {
		return nil, errors.New("failed to generate new token")
	}

	return &model.RefreshTokenResponse{
		Token: newToken,
	}, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
