package service

import (
	"2022/short-url-service/conf"
	"2022/short-url-service/dao"
	"2022/short-url-service/utils"
	"golang.org/x/sync/singleflight"
)

var (
	UrlSvc *UrlService
)

type UrlService struct {
	c            *conf.Config
	dao          *dao.Dao
	urlGenerater *utils.ShortUrlGeneratorImpl
	localCache   *utils.Lrucache
}

func New(c *conf.Config) {
	UrlSvc = &UrlService{
		c:            c,
		dao:          dao.New(c),
		urlGenerater: utils.NewShortUrlGenerator(c.UrlGenerateConf.UrlPoolSize),
		localCache:   utils.NewLrucache(c.LocalCache.Size),
	}
}

// Close ...
func Close() {
}

func (s *UrlService) GetShortUrl(lurl string) (string, error) {
	// TODO 限流 or 鉴权 验证码  异步
	surl := s.urlGenerater.GetShortKey()
	_, err := s.dao.InsertUrl(surl, lurl)
	if err != nil {
		return "", err
	}
	return surl, nil
}

var g singleflight.Group

// QueryLongLink 查询长链接
func (s *UrlService) QueryLongLink(surl string) (string, error) {
	// TODO 布隆过滤器 filter
	lurl, err := s.getLongLink(surl)
	return lurl, err
}

func (s *UrlService) getLongLink(surl string) (string, error) {
	// 查询缓存
	if v, ok := s.localCache.Get(surl); ok {
		lurl, _ := v.(string)
		return lurl, nil
	}
	// 缓存不存在，查询数据库
	lurl, err := s.loadDataFromDB(surl)
	return lurl, err
}

func (s *UrlService) loadDataFromDB(surl string) (string, error) {
	// 使用SingleFlight避免缓存击穿
	v, err, _ := g.Do(surl, func() (interface{}, error) {
		lurl, err := s.dao.SelectUrlBySurl(surl)
		if err != nil {
			return "", err
		}
		// 更新缓存
		s.localCache.Set(surl, lurl)
		return lurl, nil
	})
	return v.(string), err
}
