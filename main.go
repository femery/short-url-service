package main

import (
	"2022/short-url-service/component"
	"2022/short-url-service/conf"
	"2022/short-url-service/http"
	"2022/short-url-service/service"
)

func main() {
	if err := conf.Init("./config/application.toml"); err != nil {
		panic(err)
	}

	if err := component.InitDBByCfg(conf.Conf); err != nil {
		panic(err)
	}

	service.New(conf.Conf)
	defer service.Close()

	ginRouter := http.SetupRouter()
	panic(ginRouter.Run(":8088"))

}
