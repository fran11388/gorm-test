package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type User struct {
	Id      string
	Name    string
	Balance float64
}

type UserRepo struct {}

func (u *UserRepo) ChargeMoney(DB *gorm.DB, name string, amount int) {
	DB.Session(&gorm.Session{NewDB: false}).Exec("update user SET `balance`=`balance`+ ? where name=?", amount, name)
}

func main() {
	userRepo := UserRepo{}
	DB := GetDB()
	name := "frank"
	isErrorHappenInTransaction := true

	user := User{}
	DB.Raw("SELECT * FROM `user`\nwhere name= ?", name).Scan(&user)
	fmt.Println(fmt.Sprintf("%s 's balance is %v", name, user.Balance))

	err := DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')

		user := User{}
		tx.Raw("SELECT * FROM `user`\nwhere name= ?", name).Scan(&user)
		fmt.Println(fmt.Sprintf("in transaction , %s's balance is :%v", name, user.Balance))

		fmt.Println("charge balance now")
		userRepo.ChargeMoney(tx, name, 100)

		tx.Raw("SELECT * FROM `user`\nwhere name= ?", name).Scan(&user)
		fmt.Println(fmt.Sprintf("in transaction , %s's balance is :%v", name, user.Balance))

		if isErrorHappenInTransaction {
			// return any error will rollback
			return errors.New("some error happen :D")
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("transaction faild:%s", err.Error()))
	} else {
		fmt.Println("transaction success")
	}

	DB.Raw("SELECT * FROM `user`\nwhere name= ?", name).Scan(&user)
	fmt.Println(fmt.Sprintf("%s 's balance is %v", name, user.Balance))
}
func GetDB() *gorm.DB {
	var err error

	Username := "root"
	Password := "password"
	Address := "localhost:3306"
	Database := "gorm"

	connectInfo := fmt.Sprintf( //gcp
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		Username,
		Password,
		Address,
		Database,
	)

	db, err := gorm.Open(mysql.Open(connectInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sqldb *sql.DB
	sqldb, err = db.DB()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	maxConnections := 100
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqldb.SetMaxIdleConns(maxConnections)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqldb.SetMaxOpenConns(maxConnections)
	// SetConnMaxLifetime sets the maximum amount of timeUtil a connection may be reused.
	sqldb.SetConnMaxLifetime(time.Duration(15) * time.Minute)
	return db
}
