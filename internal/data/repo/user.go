package repo

import (
	"context"
	"kratos-layout/internal/data/query"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetById(ctx context.Context, id int64) (name string, err error) {
	mm := query.Use(r.db).User

	item, err := mm.WithContext(ctx).Where(mm.ID.Eq(id)).First()
	if err != nil {
		return
	}
	name = item.Name
	return
}
