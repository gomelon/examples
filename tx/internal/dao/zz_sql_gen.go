package dao

import (
	"context"
	"github.com/gomelon/melon/data"
	"time"
)

type UserDaoImpl struct {
	_tm *data.SQLTXManager
}

func NewUserDaoImpl(_tm *data.SQLTXManager) *UserDaoImpl {
	return &UserDaoImpl{
		_tm: _tm,
	}
}

func (_impl *UserDaoImpl) CountByBirthdayGte(ctx context.Context, time time.Time) (int32, error) {
	_sql := "select count(*) as count from `user` where birthday >= ?"
	_item := int32(0)
	_rows, _err := _impl._tm.OriginTXOrDB(ctx).
		Query(_sql, time)

	if _err != nil {
		return _item, _err
	}

	defer _rows.Close()

	if !_rows.Next() {
		return _item, _rows.Err()
	}

	_err = _rows.Scan(&_item)
	return _item, _err
}

func (_impl *UserDaoImpl) FindByBirthdayGte(ctx context.Context, time time.Time) ([]*User, error) {
	_sql := "select id, name, gender, birthday, created_at from `user` where birthday >= ?"
	_items := []*User{}
	_rows, _err := _impl._tm.OriginTXOrDB(ctx).
		Query(_sql, time)

	if _err != nil {
		return _items, _err
	}

	defer _rows.Close()

	if !_rows.Next() {
		return _items, _rows.Err()
	}

	for _rows.Next() {
		_item := &User{}
		_err = _rows.Scan(&_item.Id, &_item.Name, &_item.Gender, &_item.Birthday, &_item.CreatedAt)
		if _err != nil {
			return _items, _err
		}
		_items = append(_items, _item)
	}
	return _items, nil
}

func (_impl *UserDaoImpl) FindById(ctx context.Context, id int64) (*User, error) {
	_sql := "select id, name, gender, birthday, created_at from `user` where id = ?"
	_item := &User{}
	_rows, _err := _impl._tm.OriginTXOrDB(ctx).
		Query(_sql, id)

	if _err != nil {
		return _item, _err
	}

	defer _rows.Close()

	if !_rows.Next() {
		return _item, _rows.Err()
	}

	_err = _rows.Scan(&_item.Id, &_item.Name, &_item.Gender, &_item.Birthday, &_item.CreatedAt)
	return _item, _err
}

func (_impl *UserDaoImpl) FindByName(ctx context.Context, name string) (*User, error) {
	_sql := "SELECT id, name, gender, birthday, created_at FROM `user` WHERE (`name` = ?)"
	_item := &User{}
	_rows, _err := _impl._tm.OriginTXOrDB(ctx).
		Query(_sql, name)

	if _err != nil {
		return _item, _err
	}

	defer _rows.Close()

	if !_rows.Next() {
		return _item, _rows.Err()
	}

	_err = _rows.Scan(&_item.Id, &_item.Name, &_item.Gender, &_item.Birthday, &_item.CreatedAt)
	return _item, _err
}

func (_impl *UserDaoImpl) Insert(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO `user`(`name`,`gender`,`birthday`)" +
		"VALUES (?, ?, ?)"
	db := _impl._tm.OriginTXOrDB(ctx)
	result, err := db.Exec(query, user.Name, user.Gender, user.Birthday)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	user.Id = id
	return user, err
}

func (_impl *UserDaoImpl) UpdateById(ctx context.Context, id int64, user *User) (int64, error) {
	query := "UPDATE `user` SET `name`= ?, `gender` = ?, `birthday` = ?, `created_at` = ? " +
		"WHERE `id` = ?"
	db := _impl._tm.OriginTXOrDB(ctx)
	result, _ := db.Exec(query, user.Name, user.Gender, user.Birthday, id)
	return result.RowsAffected()
}

func (_impl *UserDaoImpl) DeleteById(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM `user` WHERE `id` = ?"
	db := _impl._tm.OriginTXOrDB(ctx)
	result, _ := db.Exec(query, id)
	return result.RowsAffected()
}
