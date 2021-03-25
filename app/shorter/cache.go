package shorter

import (
	"context"
	"strconv"
	"time"

	"github.com/n3k0fi5t/ShortTheURL/common/cache"
)

const (
	urlExpiration = 60 * time.Second
)

var (
	_   UrlDao = (*RDB)(nil)
	ctx        = context.Background()
)

type RDB cache.RDB

func convertID(id uint) string {
	return strconv.Itoa(int(id))
}

func (rdb *RDB) Get(url string) (ShortUrl, error) {
	key := make_key(url)

	str_idx, err := rdb.RDB.Get(ctx, key).Result()
	if err != nil {
		return ShortUrl{}, err
	}

	idx, err := strconv.ParseUint(str_idx, 10, 32)
	if err != nil {
		return ShortUrl{}, err
	}

	m := ShortUrl{
		ID:  uint(idx),
		URL: url,
	}
	return m, nil
}

func (rdb *RDB) GetById(id uint) (ShortUrl, error) {
	key := make_key(convertID(id))

	url, err := rdb.RDB.Get(ctx, key).Result()
	if err != nil {
		return ShortUrl{}, nil
	}

	m := ShortUrl{
		ID:  id,
		URL: url,
	}
	return m, nil
}

func (rdb *RDB) Create(urlModel *ShortUrl) (err error) {
	key := make_key(urlModel.URL)
	rkey := make_key(convertID(urlModel.ID))

	err = rdb.RDB.Set(ctx, key, urlModel.ID, urlExpiration).Err()
	if err != nil {
		return
	}

	err = rdb.RDB.Set(ctx, rkey, urlModel.URL, urlExpiration).Err()
	if err != nil {
		return
	}

	return
}

func (rdb *RDB) Delete(urlModel *ShortUrl) (err error) {
	key := make_key(urlModel.URL)

	err = rdb.RDB.Del(ctx, key).Err()
	return
}

func make_key(url string) string {
	return url
}
