package repository

import (
	"github.com/ggialluisi/nebula-back/pessoa/internal/domain/entity"
)

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
