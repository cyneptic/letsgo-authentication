package repositories

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// جایی که مستقیم به دیتابیس کار می‌کنیم

type Postgres struct {
	db *gorm.DB
}

func GormInit() (*gorm.DB, error) {
	
	host := os.Getenv("postgres_host")
	user := os.Getenv("postgres_user")
	password := os.Getenv("postgres_password")
	dbName := os.Getenv("postgres_db_name")
	port, _  := strconv.Atoi(os.Getenv("postgres_port"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(&entities.User{})
	if err != nil {
		fmt.Println(err)
	}
	return database, nil
}
func NewPostgres() *Postgres {
	db, _ := GormInit()
	return &Postgres{
		db: db,
	}
}
