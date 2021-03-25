package shorter_test

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/n3k0fi5t/ShortTheURL/app/shorter"
	"github.com/stretchr/testify/assert"
)

const (
	testAddr     = "localhost:6379"
	testPassword = ""
	testDB       = 0
	testURL      = "https://www.google.com/"
)

func TestCache(t *testing.T) {
	var rdb shorter.UrlDao
	//
	rdb_client := redis.NewClient(&redis.Options{
		Addr:     testAddr,
		Password: testPassword,
		DB:       testDB,
	})
	rdb = &shorter.RDB{RDB: rdb_client}

	_, err := rdb.Get(testURL)
	assert.Error(t, err) // key not found

	m := &shorter.ShortUrl{
		ID:  1,
		URL: testURL,
	}

	err = rdb.Create(m)
	assert.Nil(t, err)

	qm, err := rdb.Get(testURL)
	assert.Nil(t, err) // key found
	assert.Equal(t, m, &qm)

	qm, err = rdb.GetById(m.ID)
	assert.Nil(t, err) // key found
	assert.Equal(t, m, &qm)

	err = rdb.Delete(m)
	assert.Nil(t, err)

	_, err = rdb.Get(testURL)
	assert.Error(t, err) // key not found
}
