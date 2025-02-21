package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IUserService interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
}
