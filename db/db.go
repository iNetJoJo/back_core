package db

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kerlexov/back_core/model"
)

func New() *gorm.DB {
	pg := "host=" +os.Getenv("DB_HOST")+ " port=" +os.Getenv("DB_PORT")+
		  " user=" +os.Getenv("DB_USER")+ " dbname=" +os.Getenv("DB_NAME")+ " password=" +os.Getenv("DB_PASSWORD")+ ""
	db, err := gorm.Open("postgres", pg)
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	defer db.Close()

	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
	)
}

//func New() *gorm.DB {
//	db, err := gorm.Open("sqlite3", "./realworld.db")
//	if err != nil {
//		fmt.Println("storage err: ", err)
//	}
//	db.DB().SetMaxIdleConns(3)
//	db.LogMode(true)
//	return db
//}