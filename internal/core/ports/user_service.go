package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
)

type IUserService interface {
	Register(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Login(ctx context.Context, email, username, password string) (session.Session, error)
}
