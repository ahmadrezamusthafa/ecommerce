package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"github.com/google/uuid"
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

func (r *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	user.ID = uuid.NewString()
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user entity.User) (entity.User, error) {
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
