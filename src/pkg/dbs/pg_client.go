package dbs

import (
	"log"
	"os"

	"github.com/bactruongvan17/taskhub-userservice/src/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient() *gorm.DB {
	config := conf.LoadEnv().Database

	dsn := "host=" + config.PostgresHost + " port=" + config.PostgresPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname=" + config.PostgresDB + " sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect to Postgres", err)
		os.Exit(1)
	}

	log.Println("Connect to Postgres sucessfully.")
	return db
}
