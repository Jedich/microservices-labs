package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
	"user_service/utils"
)

type Database struct {
	Connection *gorm.DB
}

func OpenConnection() *Database {
	db := Database{}
	var err error
	switch s := "postgres"; s {
	case "postgres":
		retries := 3
		//var logger gorm
		l := logger.Default

		env := utils.ReadEnv("APP_DSN")
		db.Connection, err = gorm.Open(postgres.Open(env), &gorm.Config{
			Logger: l,
		})
		for err != nil {
			log.Printf("Failed to connect to database (%d)\n", retries)
			if retries > 1 {
				retries--
				time.Sleep(5 * time.Second)
				db.Connection, err = gorm.Open(postgres.Open(env), &gorm.Config{
					Logger: l,
				})
				continue
			}
			log.Fatal(err.Error())
		}
	default:
		log.Fatal(fmt.Sprintf("Unsupported driver %s", s))
	}
	log.Println("Connected to database")

	log.Println("Running migrations")
	m, err := migrate.New("file://migrations", utils.ReadEnv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Drop(); err != nil {
		panic(err)
	}

	m, err = migrate.New("file://migrations", utils.ReadEnv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		panic(err)
	}

	return &db
}

func (db *Database) CloseConnection() {
	sqlDB, err := db.Connection.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDB.Close()
}
