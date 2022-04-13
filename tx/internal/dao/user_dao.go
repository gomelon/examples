package dao

import (
	"context"
	"github.com/gomelon/examples/tx/internal/entity"
	"time"
)

//UserDao
//+melon.sql.Mapper `TableName:"user"`
type UserDao interface {
	FindById(ctx context.Context, id int64) *entity.User
	FindByName(ctx context.Context, name string) *entity.User
	FindByBirthdayGte(ctx context.Context, time time.Time) []*entity.User
	Insert(ctx context.Context, user *entity.User) *entity.User
	UpdateById(ctx context.Context, id int64, user *entity.User) int64
	DeleteById(ctx context.Context, id int64) int64
}
