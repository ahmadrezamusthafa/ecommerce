package model

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"time"
)

type UserTable struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"password" json:"-"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func (UserTable) TableName() string {
	return "users"
}

func (t UserTable) ToEntity() entity.User {
	return entity.User{
		ID:        t.ID,
		Name:      t.Name,
		Username:  t.Username,
		Email:     t.Email,
		Password:  t.Password,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
