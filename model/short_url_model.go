package model

const DefaultLink = "https://www.baidu.com/"

const DomainNamePrefix = "http://localhost:8088/"
const DomainNameMiddle = "surl/v/"

type LongUrlParams struct {
	Url string `json:"url"`
	Ttl int64  `json:"ttl"`
}

const LUrlLongLimit = 750
