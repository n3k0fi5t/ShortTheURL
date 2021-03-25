package dbx

import (
	"fmt"

	"github.com/n3k0fi5t/ShortTheURL/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB encapsulate the actual ORM
type DB struct {
	Orm *gorm.DB
}

const (
	dsnFormat = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func Init(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(dsnFormat, cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)
	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{orm}, nil
}

func (db *DB) Close() (err error) {
	// gorm does not have close method, just leave thing here
	return
}
