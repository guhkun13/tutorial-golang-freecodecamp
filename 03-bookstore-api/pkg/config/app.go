package config

// import "gorm.io/driver/mysql"
// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	
	d, err := gorm.Open("mysql", "root:root123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
	
}

func GetDB() *gorm.DB {
	return db
}