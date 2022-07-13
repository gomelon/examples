package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomelon/melon/data"
	"testing"
	"time"
)

func TestCRUD(t *testing.T) {
	// prepare
	tm, closeFunc := tm()
	defer closeFunc()
	var userDao UserDao = NewUserDaoImpl(tm)

	// execute
	user := &User{
		Name:     "GoMelon",
		Gender:   0,
		Birthday: time.Now(),
	}
	ctx := context.Background()
	user, err := userDao.Insert(ctx, user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert User Id:", user.Id, ",Name:", user.Name)

	user, err = userDao.FindById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

	if user == nil {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("Find User Id:", user.Id, ",Name:", user.Name)
	}

	// clean
	_, err = userDao.DeleteById(ctx, user.Id)
	if err != nil {
		panic(err)
	}
}

func TestTransactionRollback(t *testing.T) {
	// prepare
	tm, closeFunc := tm()
	defer closeFunc()
	var userDao UserDao = NewUserDaoImpl(tm)

	// execute
	ctx := context.Background()
	newCtx, err := tm.Begin(ctx, nil)
	if err != nil {
		panic(err)
	}
	user := &User{
		Name:     "GoMelon",
		Gender:   0,
		Birthday: time.Now(),
	}
	user, err = userDao.Insert(newCtx, user)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return
	}
	tm.Rollback(newCtx)

	user, err = userDao.FindById(ctx, user.Id)
	if err != nil {
		panic(err)
	}

	if user == nil {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("Find User ID:", user.Id, ",Name:", user.Name)
		userDao.DeleteById(ctx, user.Id)
	}
}

func tm() (tm *data.SQLTXManager, closeFunc func()) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/user?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	tm = data.NewSqlTxManager("user", db)
	closeFunc = func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}
	return
}
