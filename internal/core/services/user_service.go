package services

import (
	"context"
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository ports.IUserRepository
	sessionCfg     *session.Config
}

func NewUserService(sessionCfg *session.Config, userRepository ports.IUserRepository) ports.IUserService {
	return &userService{
		sessionCfg:     sessionCfg,
		userRepository: userRepository,
	}
}

func (s *userService) Register(ctx context.Context, user entity.User) (entity.User, error) {
	if err := s.validateUniqueUser(ctx, user); err != nil {
		return entity.User{}, err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = hashedPassword

	return s.userRepository.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, user entity.User) (entity.User, error) {
	if err := s.validateUniqueUser(ctx, user); err != nil {
		return entity.User{}, err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = hashedPassword

	return s.userRepository.Update(ctx, user)
}

func (s *userService) Login(ctx context.Context, email, username, password string) (session.Session, error) {
	var (
		user entity.User
		err  error
	)

	if email != "" {
		user, err = s.userRepository.GetByEmail(ctx, email)
		if err != nil {
			return session.Session{}, err
		}
	} else if username != "" {
		user, err = s.userRepository.GetByUsername(ctx, username)
		if err != nil {
			return session.Session{}, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return session.Session{}, errors.New("invalid credentials")
	}

	token, err := s.sessionCfg.GenerateToken(user.ID)
	if err != nil {
		return session.Session{}, err
	}

	return token, nil
}

func (s *userService) validateUniqueUser(ctx context.Context, user entity.User) error {
	if _, err := s.userRepository.GetByEmail(ctx, user.Email); err == nil && user.ID == 0 {
		return errors.New("email already in use")
	}
	if _, err := s.userRepository.GetByUsername(ctx, user.Username); err == nil && user.ID == 0 {
		return errors.New("username already in use")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
