package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)


func NewDatabase() (*gorm.DB, error){
	fmt.Println("Setting up new database connection")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_Password")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
    fmt.Println(dbUsername)
    fmt.Println(dbPassword)
    fmt.Println(dbHost)
    fmt.Println(dbTable)
    fmt.Println(dbPort)
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort,dbUsername, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectString)
	if err != nil{
		return db,err
	}
	if err := db.DB().Ping(); err!=nil{
		return db, err
	}
	return db,nil
}