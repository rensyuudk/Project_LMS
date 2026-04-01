package database

import (
	"fmt"
	"lms/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:2224@tcp(127.0.0.1:3306)/lms_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("detail error: ", err)
		panic("gagal tersambung ke database: ")
	}
	fmt.Println("berhasil koneksi")

	err = db.AutoMigrate(
		&models.Student{},
		&models.Course{},
		&models.Module{},
		&models.Enrollment{},
		&models.Score{},
	)
	if err != nil {
		fmt.Println("detail error: ", err)
		panic("gagal migarsi database: ")
	}
	fmt.Println("migrasi berhasil")

	DB = db
}
