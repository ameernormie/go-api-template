package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func GetDbInfo() (User, Password, Host, Port, Database string) {
	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Database = os.Getenv("DB_DATABASE")

	fmt.Println("Host: ", Host)
	fmt.Println("Port: ", Port)
	fmt.Println("User: ", User)
	fmt.Println("Password: ", Password)
	fmt.Println("Database: ", Database)

	if Host == "" || Port == "" || User == "" || Password == "" || Database == "" {
		panic("Missing fields in env")
	}
	return
}

// Connect : connect database
func Connect() (db *gorm.DB, err error) {
	User, Password, Host, Port, Database := GetDbInfo()

	connectString := fmt.Sprintf(
		`host=%s port=%s user=%s dbname=%s password=%s sslmode=disable`,
		Host,
		Port,
		User,
		Database,
		Password,
	)

	db, err = gorm.Open("postgres", connectString)
	return
}
