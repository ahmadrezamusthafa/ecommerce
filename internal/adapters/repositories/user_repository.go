package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	result := r.db.WithContext(ctx).Omit("created_at").Save(&user)
	if result.Error != nil {
		return entity.User{}, fmt.Errorf("failed to update user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var user entity.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}
