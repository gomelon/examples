package impl

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomelon/examples/tx/internal/dao"
	"github.com/gomelon/examples/tx/internal/entity"
	"github.com/gomelon/melon"
	"testing"
	"time"
)

func TestCRUD(t *testing.T) {
	// prepare
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/example?charset=utf8&parseTime=True")
	melon.PanicOnError(err)
	defer func(db *sql.DB) {
		err := db.Close()
		melon.PanicOnError(err)
	}(db)

	dbProvider := melon.NewDBProvider(melon.DBNameDefault, db, nil)
	melon.RegisterDBProvider(dbProvider)

	var userDao dao.UserDao = &UserDaoImpl{}

	// execute
	user := &entity.User{
		Name:     "GoMelon",
		Gender:   0,
		Birthday: time.Now(),
	}
	ctx := context.Background()
	userDao.Insert(ctx, user)

	user = userDao.FindById(ctx, user.ID)

	if user == nil {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("ID:", user.ID, ",Name:", user.Name)
	}

	// clean
	userDao.DeleteById(ctx, user.ID)
}

func TestTransactionRollback(t *testing.T) {
	// prepare
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/example?charset=utf8&parseTime=True")
	melon.PanicOnError(err)
	defer func(db *sql.DB) {
		err := db.Close()
		melon.PanicOnError(err)
	}(db)

	dbProvider := melon.NewDBProvider(melon.DBNameDefault, db, &melon.SqlTxManager{})
	melon.RegisterDBProvider(dbProvider)

	var userDao dao.UserDao = &UserDaoImpl{}

	// execute
	dbProvider = melon.GetDBProvider(melon.DBNameDefault)
	ctx := context.Background()
	newCtx, err := dbProvider.Begin(ctx, nil)
	user := &entity.User{
		Name:     "GoMelon",
		Gender:   0,
		Birthday: time.Now(),
	}
	userDao.Insert(newCtx, user)
	dbProvider.Rollback(newCtx)

	user = userDao.FindById(ctx, user.ID)

	if user == nil {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("ID:", user.ID, ",Name:", user.Name)
		userDao.DeleteById(ctx, user.ID)
	}
}
