package db

import (
	"fmt"

	"github.com/minhquy1903/online-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	cfg.Postgres.PostgresHost,
	cfg.Postgres.PostgresPort,
	cfg.Postgres.PostgresUser,
	cfg.Postgres.PostgresDbName,
	cfg.Postgres.PostgresPassword,
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}