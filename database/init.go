package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type File struct {
	ID          uint `gorm:"primaryKey;autoIncrement;not null"`
	Title       string
	Description string
	FilePath    string
}

var Db *gorm.DB

func Init() {
	godotenv.Load(".env")
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&File{})
	if err != nil {
		panic("failed to migrate database schema")
	}

	if !db.Migrator().HasTable(&File{}) {
		err = db.Migrator().CreateTable(&File{})
		if err != nil {
			panic("failed to create table File")
		}
		fmt.Println("Table created.")
	} else {
		fmt.Println("Table arleady exists")
	}
	Db = db
}
