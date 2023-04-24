package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

var DB *sql.DB

func Init() {

	var err error
	err = godotenv.Load(".env")

	var configDB DatabaseConfig
	configDB.Host = os.Getenv("DB_HOST")
	configDB.Port = os.Getenv("DB_PORT")
	configDB.User = os.Getenv("DB_USERNAME")
	configDB.Pass = os.Getenv("DB_PASSWORD")
	configDB.Name = os.Getenv("DB_NAME")

	pg_con_string := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", configDB.Host, configDB.Port, configDB.User, configDB.Pass, configDB.Name)

	DB, err = sql.Open("postgres", pg_con_string)

	if err != nil {
		fmt.Println(err)
	}

	if err = DB.Ping(); err != nil {
		fmt.Println(err)
	}

}
