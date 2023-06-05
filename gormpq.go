package pqgorm // this is any name ...this is not related to sqlx

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Struct to parse, pass and get data easily.
type User struct {
	ID       uint   `gorm:"primary_key"` // using GORM can define keys, null, notNUll constraints
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}

func connectToDatabase() (*gorm.DB, error) {
	host := "localhost"
	port := 5432
	user := "pipo"     // use your username
	pwd := "mypwdtest" // use your password , pass it securely not like this
	dbName := "practice"

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbName,
		pwd,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&User{})

	return db, err
}
