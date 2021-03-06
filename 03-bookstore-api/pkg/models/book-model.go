package models

import (
  "fmt"
	"github.com/jinzhu/gorm"
	"github.com/guhkun13/tutorial/freeCodeCamp/03-bookstore-api/pkg/config"
)

var db *gorm.DB 

type Book struct {
	gorm.Model 
	Name string `gorm:"json":"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(&b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book 
	if db := db.Where("ID=?", Id).First(&getBook); db.Error != nil {
    fmt.Println("object not found!!!")
  }
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
  return book
}