package config

import (
	"backend-absensi/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// Connect ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal koneksi ke database: ", err)
	}

	// Simpan ke variabel global
	DB = db

	// Migrate model
	err = db.AutoMigrate(
		&model.Departement{},
		&model.Employee{},
		&model.Attendance{},
		&model.AttendanceHistory{},
	)

	db.Exec("ALTER TABLE departements MODIFY COLUMN max_clock_in_time TIME")
	db.Exec("ALTER TABLE departements MODIFY COLUMN max_clock_out_time TIME")
	if err != nil {
		log.Fatal("❌ Gagal migrate model: ", err)
	}

	fmt.Println("✅ Database connected & migrated")
}
