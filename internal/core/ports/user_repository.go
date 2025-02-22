package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}
