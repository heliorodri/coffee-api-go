package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gorm.io/driver/postgres"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type DBConnection struct{}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

type Config struct {
	Database DatabaseConfig `toml:"database"`
}

func ReadConfigFile() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var config Config

	fileContent, err := ioutil.ReadFile("config/database.toml")

	if err != nil {
		return config
	}

	dbConfig := fmt.Sprintf(string(fileContent), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	println("dbConfig: ", dbConfig)

	if _, err := toml.Decode(string(dbConfig), &config); err != nil {
		panic(err)
	}

	println("Database config: ")
	println("host: ", config.Database.Host)
	println("port: ", config.Database.Port)
	println("user: ", config.Database.User)
	println("password: ", config.Database.Password)
	println("dbname: ", config.Database.DBName)

	return config
}

func NewConnection() *gorm.DB {
	cfg := ReadConfigFile()

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, strconv.Itoa(cfg.Database.Port), cfg.Database.User, cfg.Database.Password, cfg.Database.DBName,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil
	}

	return db
}

func NewDBModule() fx.Option {
	return fx.Provide(
		func() Config {
			return ReadConfigFile()
		},
		NewConnection,
	)
}