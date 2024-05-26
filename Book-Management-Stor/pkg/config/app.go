package config

//import gorm package
import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// create variable db so that every folder can interect Databse through this db

var (
	db *gorm.DB
)

// Function Connect which will connect db with database

func Connect() {
	fmt.Println("")
	d, err := gorm.Open("mysql", "swarnendu:Swarna@2003/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// Function GetDB return db
func GetDB() *gorm.DB {
	return db
}
