package shorter_test

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/n3k0fi5t/ShortTheURL/app/shorter"
	"github.com/n3k0fi5t/ShortTheURL/common/cache"
	"github.com/n3k0fi5t/ShortTheURL/common/dbx"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	testAddr     = "localhost:6379"
	testPassword = ""
	testDB       = 0
	testDSN      = "host=localhost user=user password=password dbname=db port=5432 sslmode=disable"
	testURL      = "https://www.google.com/"
)

func TestService(t *testing.T) {
	//init
	orm, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	assert.Nil(t, err)

	orm.AutoMigrate(&shorter.ShortUrl{})
	db := &dbx.DB{orm}

	rdb_client := redis.NewClient(&redis.Options{
		Addr:     testAddr,
		Password: testPassword,
		DB:       testDB,
	})
	rdb := &cache.RDB{RDB: rdb_client}

	serv := shorter.NewService(8, db, rdb)

	surl, err := serv.Shorten(testURL)
	assert.Nil(t, err)
	fmt.Println(surl)

	origin_url, err := serv.Proxy(surl)
	assert.Nil(t, err)
	assert.Equal(t, testURL, origin_url)
}
