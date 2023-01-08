package repo

import (
	"context"
	"github.com/ebar-go/ego/component"
	"github.com/ebar-go/ego/component/db"
	"unknown/internal/domain/entity"
)

type UserRepo struct {
	database *db.Instance
}

func NewUserRepo() *UserRepo {
	return &UserRepo{database: component.DB()}
}

func (repo *UserRepo) Create(ctx context.Context, item *entity.UserEntity) error {
	return repo.database.WithContext(ctx).Where("phone = ?", item.Phone).FirstOrCreate(item).Error
}
