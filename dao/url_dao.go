package dao

import (
	"2022/short-url-service/component"
	"2022/short-url-service/conf"
	"github.com/jinzhu/gorm"
	"time"
)

type ShortURLs struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(16);not null"`
	Surl      string    `gorm:"unique;column:surl;type:varchar(255)"`                               // 短链标志
	Lurl      string    `gorm:"column:lurl;type:varchar(255)"`                                      // 原始链接
	Ctime     time.Time `gorm:"column:ctime;type:datetime;not null;default:CURRENT_TIMESTAMP"`      // 创建时间
	LastVtime time.Time `gorm:"column:last_vtime;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 最后访问时间
	Times     int       `gorm:"column:times;type:int(16);not null;default:1"`                       // 访问次数
	Status    int       `gorm:"column:status;type:int(2);not null;default:0"`                       // 0正常；1失效
}

// Dao dao.
type Dao struct {
	db *gorm.DB
}

func (ShortURLs) TableName() string {
	return "short_urls"
}

func New(c *conf.Config) *Dao {
	return &Dao{
		db: component.DB,
	}
}

func (d *Dao) InsertUrl(surl string, lurl string) (*ShortURLs, error) {
	url := &ShortURLs{
		Lurl: lurl,
		Surl: surl,
	}
	if err := d.db.Create(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

func (d *Dao) DeleteUrl(surl string) error {
	if err := d.db.Where(&ShortURLs{Surl: surl}).Delete(ShortURLs{}).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) SelectUrlByLUrl(lurl string) (*ShortURLs, error) {
	url := &ShortURLs{}
	err := d.db.Model(&ShortURLs{}).Select("id").Where(&ShortURLs{Lurl: lurl}).Find(url).Error
	if err != nil {
		return nil, err
	}
	return url, err
}

func (d *Dao) SelectUrlBySurl(surl string) (string, error) {
	url := &ShortURLs{}
	err := d.db.Model(&ShortURLs{}).Select("lurl").Where(&ShortURLs{Surl: surl}).Find(url).Error
	if err != nil {
		return "", err
	}
	return url.Lurl, err
}
