package core

import (
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type AppConfig struct {
	Port int
	Db   DbConfig
}

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

var Config AppConfig

var appConfigOnce sync.Once

func GetAppConfig() *AppConfig {
	appConfigOnce.Do(func() {
		log.Println("Initializing app config from env")
		err := envconfig.Process("app", &Config)

		if err != nil {
			log.Fatalf("Error parsing env for app config: %v", err)
		}
	})

	return &Config
}

const DbDriver = "postgres"

var db *sql.DB
var dbConfigOnce sync.Once

func GetDb() *sql.DB {
	dbConfigOnce.Do(func() {
		dbConfig := GetAppConfig().Db
		connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Dbname)

		newDb, err := sql.Open(DbDriver, connectionString)
		db = newDb
		if err != nil {
			log.Fatalf("Error establishing connection to db: %v", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("Error pinging db %v", err)
		}
	})

	return db
}
