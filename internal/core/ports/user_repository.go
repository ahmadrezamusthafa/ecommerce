package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
}
