package config

import (
	"fmt"

	// "github.com/go-sql-driver/mysql"
	// "gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:Justdial@123@/bookstore?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:Justdial@123@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local")
	// dsn := "root:Justdial@123@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB", db)
	db = d
}

// func Connect() {
//     dsn := "root:Justdial@123@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
//     d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database: " + err.Error())
//     }
//     db = d
//     fmt.Println("Connected to DB")
// }

func GetDB() *gorm.DB {
	return db
}

// package config

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var (
// 	db *gorm.DB
// )

// func Connect() {
// 	dsn := "root:Justdial@123@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database: " + err.Error())
// 	}
// 	db = d
// 	fmt.Println("Connected to DB", d)
// }

// func GetDB() *gorm.DB {
// 	return db
// }
