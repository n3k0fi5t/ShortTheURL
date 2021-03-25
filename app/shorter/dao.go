package shorter

// UrlDao persists url in database and cache
type UrlDao interface {
	Get(url string) (ShortUrl, error)
	GetById(id uint) (ShortUrl, error)
	Create(urlModel *ShortUrl) error
	Delete(urlModel *ShortUrl) error
}
