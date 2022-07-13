package dao

import (
	"context"
	"time"
)

//UserDao 用户信息Dao
//sql:table name=`user` dialect="mysql"
type UserDao interface {
	//FindById 通过ID获取用户信息
	/*sql:select query="select * from `user` where id = :id" master*/
	FindById(ctx context.Context, id int64) (*User, error)
	/*sql:select query="select * from `user` where birthday >= :time"*/
	FindByBirthdayGte(ctx context.Context /*sql:param ctx*/, time time.Time) ([]*User, error)
	/*sql:select query="select count(*) as count from `user` where birthday >= :time"*/
	CountByBirthdayGte(ctx context.Context /*sql:param ctx*/, time time.Time) (int32, error)
	//FindByName 通过用户名获取用户信息
	FindByName(ctx context.Context, name string) (*User, error)
	Insert(ctx context.Context, user *User) (*User, error)
	UpdateById(ctx context.Context, id int64, user *User) (int64, error)
	DeleteById(ctx context.Context, id int64) (int64, error)
}
