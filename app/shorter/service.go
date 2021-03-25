package shorter

import (
	"errors"

	"github.com/n3k0fi5t/ShortTheURL/app/shorter/base62"
	"github.com/n3k0fi5t/ShortTheURL/common/cache"
	"github.com/n3k0fi5t/ShortTheURL/common/dbx"
)

var (
	_                Service = &service{}
	ErrProxyNotFound         = errors.New("Proxy not found")
)

type Service interface {
	Shorten(url string) (string, error)
	Proxy(url string) (string, error)
}

type service struct {
	sUrlLength int
	DB         UrlDao
	RDB        UrlDao
}

func NewService(len int, db *dbx.DB, rdb *cache.RDB) service {
	return service{len, (*DB)(db), (*RDB)(rdb)}
}

func (serv *service) Shorten(url string) (string, error) {
	var surl ShortUrl
	// 1) query from cache, return directly if hit
	if surl, err := serv.RDB.Get(url); err == nil {
		id := surl.ID
		pattern, err := base62.Encode(id, serv.sUrlLength)
		return pattern, err
	}

	// 2) create and insert to db
	surl = ShortUrl{URL: url}
	if err := serv.DB.Create(&surl); err != nil {
		return "", err
	}

	// 3) keep a copy to cache
	serv.RDB.Create(&surl)

	pattern, err := base62.Encode(surl.ID, serv.sUrlLength)
	return pattern, err
}

func (serv *service) Proxy(encoded string) (string, error) {
	id, err := base62.Decode(encoded)
	if err != nil {
		return "", err
	}

	// try get from cache first
	model, err := serv.RDB.GetById(id)
	if err == nil {
		return model.URL, nil
	}

	model, err = serv.DB.GetById(id)
	if err != nil {
		return "", ErrProxyNotFound
	}

	return model.URL, err
}
