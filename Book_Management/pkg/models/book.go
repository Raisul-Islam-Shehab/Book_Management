package models

import(
	"gorm.io/gorm"
	"Book_Management/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}