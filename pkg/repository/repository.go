package repository

import (
	"log"
	"time"

	"github.com/pirosiki197/sodan-grpc/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	SodanRepository
	ReplyRepository
	AutoMigrate(value any) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(config *config.Config) Repository {
	db, err := connectDB(config)
	if err != nil {
		log.Fatal(err)
	}

	return &repository{db: db}
}

func connectDB(config *config.Config) (*gorm.DB, error) {
	for i := 0; i < 10; i++ {
		db, err := gorm.Open(mysql.Open(config.DBConf.FormatDSN()), &gorm.Config{})
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		return db, nil
	}
	return nil, nil
}

func (r *repository) AutoMigrate(value any) error {
	return r.db.AutoMigrate(value)
}
