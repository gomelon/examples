package impl

import (
	"context"
	"github.com/gomelon/examples/tx/internal/entity"
	"github.com/gomelon/melon"
	"time"
)

type UserDaoImpl struct {
}

func (u UserDaoImpl) FindById(ctx context.Context, id int64) *entity.User {
	query := "SELECT `id`, `name`, `gender`, `birthday`, `created_at` " +
		"FROM `user` " +
		"WHERE `id` = ?"
	db := melon.GetSqlExecutor(ctx, melon.DBNameDefault)
	rows, err := db.Query(query, id)
	melon.PanicOnError(err)
	defer rows.Close()
	if rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Birthday, &user.CreatedAt)
		melon.PanicOnError(err)
		return &user
	} else {
		return nil
	}
}

func (u UserDaoImpl) FindByName(ctx context.Context, name string) *entity.User {
	panic("implement me")
}

func (u UserDaoImpl) FindByBirthdayGte(ctx context.Context, time time.Time) []*entity.User {
	query := "SELECT `id`, `name`, `gender`, `birthday`, `created_at` " +
		"FROM `user` " +
		"WHERE `birthday` >= ?"
	db := melon.GetSqlExecutor(ctx, melon.DBNameDefault)
	rows, err := db.Query(query, time)
	melon.PanicOnError(err)
	defer rows.Close()
	var result = make([]*entity.User, 0, 2)
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Birthday, &user.CreatedAt)
		melon.PanicOnError(err)
		result = append(result, &user)
	}
	return result
}

func (u UserDaoImpl) Insert(ctx context.Context, user *entity.User) *entity.User {
	query := "INSERT INTO `user`(`name`,`gender`,`birthday`)" +
		"VALUES (?, ?, ?)"
	db := melon.GetSqlExecutor(ctx, melon.DBNameDefault)
	result, err := db.Exec(query, user.Name, user.Gender, user.Birthday)
	melon.PanicOnError(err)
	id, err := result.LastInsertId()
	melon.PanicOnError(err)
	user.ID = id
	return user
}

func (u UserDaoImpl) UpdateById(ctx context.Context, id int64, user *entity.User) int64 {
	query := "UPDATE `user` SET `name`= ?, `gender` = ?, `birthday` = ?, `created_at` = ? " +
		"WHERE `id` = ?"
	db := melon.GetSqlExecutor(ctx, melon.DBNameDefault)
	result, err := db.Exec(query, user.Name, user.Gender, user.Birthday, id)
	melon.PanicOnError(err)
	rowsAffected, err := result.RowsAffected()
	melon.PanicOnError(err)
	return rowsAffected
}

func (u UserDaoImpl) DeleteById(ctx context.Context, id int64) int64 {
	query := "DELETE FROM `user` WHERE `id` = ?"
	db := melon.GetSqlExecutor(ctx, melon.DBNameDefault)
	result, err := db.Exec(query, id)
	melon.PanicOnError(err)
	rowsAffected, err := result.RowsAffected()
	melon.PanicOnError(err)
	return rowsAffected
}
