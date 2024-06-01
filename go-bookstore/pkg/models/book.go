package models

import (
	"fmt"
	"go-bookstore/pkg/config"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
	}
}

func (b *Book) CreateBook() *Book {
	fmt.Println("creating........", b)
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// func (b *Book) CreateBook() *Book {
// 	fmt.Println("creating........")
// 	if err := db.Create(b).Error; err != nil {
// 		fmt.Println("Error creating book:", err)
// 		return nil
// 	}
// 	return b
// }

// func GetAllBooks() []Book {
// 	fmt.Println("getinggggg...")
// 	var Books []Book
// 	db.Find(&Books)
// 	fmt.Println("bbbbbooooookkkkksss:", Books)
// 	return Books
// }

func GetAllBooks() []Book {
	fmt.Println("gettinggggg...")
	var Books []Book
	if err := db.Find(&Books).Error; err != nil {
		fmt.Println("Error getting books:", err)
		return nil
	}
	fmt.Println("bbbbbooooookkkkksss:", Books)
	return Books
}

//	func GetBookById(Id int64) (*Book, *gorm.DB) {
//		var getBook Book
//		db := db.Where("ID=?", Id).Find(&getBook)
//		return &getBook, db
//	}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID = ?", Id).First(&getBook)
	if result.Error != nil {
		fmt.Println("Error getting book by ID:", result.Error)
		return nil, result
	}
	return &getBook, result
}

//	func DeleteBookById(Id int64) Book {
//		var book Book
//		db.Where("ID=?", Id).Delete(book)
//		return book
//	}
func DeleteBookById(Id int64) error {
	var book Book
	if err := db.Where("ID = ?", Id).Delete(&book).Error; err != nil {
		fmt.Println("Error deleting book:", err)
		return err
	}
	return nil
}
