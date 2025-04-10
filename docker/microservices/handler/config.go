package handler

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func InfoFromEnv() *gorm.DB {

	dbUserName := os.Getenv("POSTGRES_USER")
	dbUserPwd := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_INTERNAL_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	// dbUserName := "vibav"
	// dbUserPwd := "vibav"
	// dbHost := "database"
	// dbPort := "5432"
	// dbName := "db"
	dbSchema := "public"
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUserName, dbUserPwd,
		dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbSchema, // schema name
			SingularTable: false,
		},
	})

	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
 
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}




