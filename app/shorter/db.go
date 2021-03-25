package shorter

import (
	"github.com/n3k0fi5t/ShortTheURL/common/dbx"
)

type DB dbx.DB

var (
	_ UrlDao = (*DB)(nil)
)

func (db *DB) Get(url string) (ShortUrl, error) {
	var model ShortUrl

	tx := db.Orm.Begin()
	if err := tx.Where("url = ?", url).Take(&model).Error; err != nil {
		tx.Rollback()
		return ShortUrl{}, err
	}

	err := tx.Commit().Error
	return model, err
}

func (db *DB) GetById(id uint) (ShortUrl, error) {
	var model ShortUrl

	tx := db.Orm.Begin()
	if err := tx.Where("id = ?", id).Take(&model).Error; err != nil {
		tx.Rollback()
		return ShortUrl{}, err
	}

	err := tx.Commit().Error
	return model, err
}

func (db *DB) Create(urlModel *ShortUrl) (err error) {
	err = db.Orm.Create(urlModel).Error
	return
}

func (db *DB) Delete(urlModel *ShortUrl) (err error) {
	err = db.Orm.Where(urlModel).Delete(ShortUrl{}).Error
	return
}
