package shorter_test

import (
	"testing"

	"github.com/n3k0fi5t/ShortTheURL/app/shorter"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	testDSN = "host=localhost user=user password=password dbname=db port=5432 sslmode=disable"
	testURL = "https://www.google.com/"
)

func TestDB(t *testing.T) {
	var db shorter.UrlDao
	orm, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	assert.Nil(t, err)

	orm.AutoMigrate(&shorter.ShortUrl{})

	db = &shorter.DB{orm}

	m := &shorter.ShortUrl{URL: testURL}

	_, err = db.Get(testURL) // cannot get
	assert.Error(t, err)

	err = db.Create(m)
	assert.Nil(t, err)

	rm, err := db.Get(testURL)
	assert.Nil(t, err)
	assert.Equal(t, rm.URL, m.URL)

	rm2, err := db.GetById(rm.ID)
	assert.Nil(t, err)
	assert.Equal(t, rm2.URL, m.URL)

	err = db.Delete(&rm)
	assert.Nil(t, err)

	_, err = db.Get(testURL) // cannot get
	assert.Error(t, err)
}
