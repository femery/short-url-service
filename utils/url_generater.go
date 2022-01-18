package utils

import "github.com/teris-io/shortid"

type ShortUrlGenerator interface {
	genUrlWorker()
	GetShortKey() string
}

type ShortUrlGeneratorImpl struct {
	urlKeyPool chan string
	sid        *shortid.Shortid
}

func NewShortUrlGenerator(poolSize int32) *ShortUrlGeneratorImpl {
	urlPool := make(chan string, poolSize)
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic("over")
	}
	entity := &ShortUrlGeneratorImpl{
		urlKeyPool: urlPool,
		sid:        sid,
	}
	entity.genUrlWorker()
	return entity
}

func (s *ShortUrlGeneratorImpl) genUrlWorker() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
		// start
		for {
			urlId, err := s.sid.Generate()
			if err != nil {
				return
			}
			s.urlKeyPool <- urlId
		}
		// end
	}()
}

func (s *ShortUrlGeneratorImpl) GetShortKey() string {
	return <-s.urlKeyPool
}
