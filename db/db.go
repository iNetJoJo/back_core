package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kerlexov/back_core/model"
)

func New() *gorm.DB {
	//pg := "host=" +os.Getenv("DB_HOST")+ " port=" +os.Getenv("DB_PORT")+ " user=" +os.Getenv("DB_USER")+ " dbname=" +os.Getenv("DB_NAME")+ " password=" +os.Getenv("DB_PASSWORD")+ ""
	db, err := gorm.Open("postgres", "host=db port=5432 user=stivi dbname=baza password=popivi sslmode=disable")
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	return db
}
func Close(db *gorm.DB) {
	defer db.Close()
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