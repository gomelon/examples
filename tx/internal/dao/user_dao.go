package dao

import (
	"context"
	"time"
)

//UserDao 用户信息Dao
//+sqlmap.Mapper Table=`user` Dialect="mysql"
type UserDao interface {
	//FindById 通过ID获取用户信息
	/*+sqlmap.Select Query="select * from `user` where id = :id" Master*/
	FindById(ctx context.Context, id int64) (*User, error)
	/*+sqlmap.Select Query="select * from `user` where birthday >= :time"*/
	FindByBirthdayGte(ctx context.Context /*sql:param ctx*/, time time.Time) ([]*User, error)
	/*+sqlmap.Select Query="select count(*) as count from `user` where birthday >= :time"*/
	CountByBirthdayGte(ctx context.Context /*sql:param ctx*/, time time.Time) (int32, error)
	//FindByName 通过用户名获取用户信息
	FindByName(ctx context.Context, name string) (*User, error)
	Insert(ctx context.Context, user *User) (*User, error)
	UpdateById(ctx context.Context, id int64, user *User) (int64, error)
	DeleteById(ctx context.Context, id int64) (int64, error)
}
