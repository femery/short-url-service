package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

var (
	Conf = &Config{}
)

type Config struct {
	MySQLConf       *MySQLConf `toml:"MySQLConf"`
	UrlGenerateConf *UrlGenerateConf
	LocalCache      *LocalCache
}

type MySQLConf struct {
	Host     string `toml:"Host"`
	Port     string
	User     string
	Password string
	DbName   string
}

type UrlGenerateConf struct {
	UrlPoolSize int32
	ReTry       bool
}

type LocalCache struct {
	Open bool
	Size int32
	ttl  int32
}

// Init conf.
func Init(filePath string) (err error) {
	return ReadConf(filePath)
}

func ReadConf(filePath string) (err error) {
	var (
		fp       *os.File
		fcontent []byte
	)
	if fp, err = os.Open(filePath); err != nil {
		return
	}

	if fcontent, err = ioutil.ReadAll(fp); err != nil {
		return
	}

	if err = toml.Unmarshal(fcontent, Conf); err != nil {
		return
	}
	fmt.Println(Conf)
	return
}
