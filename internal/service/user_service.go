package service

import (
	"database/sql"
	"github.com/dian/bank-api/internal/db"
	"github.com/dian/bank-api/internal/model"
	"github.com/dian/bank-api/internal/repository"
	"github.com/dian/bank-api/pkg/errors"
	"github.com/dian/bank-api/pkg/jwtutil"
	"time"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}
func CreateUser(name, email string) error {
	return repository.InsertUser(db.DB, name, email)
}

func ListUsers() ([]model.User, error) {
	return repository.GetAllUsers(db.DB)
}

func (s *UserService) GetUserByEmail(email string) (*model.UserResponseDTO, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewBusinessError("user not found")
		}
		return nil, errors.NewInternalError("-Err- failed to get user by email", err)
	}

	dto := model.ToUserResponseDTO(*user)
	return &dto, nil
}

func (s *UserService) Login(email string) (*model.LoginResponse, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	token, expiredAt, err := jwtutil.GenerateToken(user.ID, user.Email, "user", 3*time.Minute)
	if err != nil {
		return nil, err
	}

	// Simpan ke sessions

	err = s.repo.SaveSession(user.ID, token, expiredAt)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     token,
		ExpiredAt: expiredAt.Unix(),
	}, nil
}
