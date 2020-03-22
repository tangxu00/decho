package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Wuvist/decho/conf"
	"github.com/Wuvist/decho/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	webApp, err := getWebApp()
	if err != nil {
		panic(err)
	}
	webApp.Run()
}

// WebApp represent a web application
type WebApp struct {
	*echo.Echo
	config *conf.Config
	blog   *controller.BlogController
}

// Run the web app
func (e *WebApp) Run() {
	e.Logger.Fatal(e.Start(":1323"))
}

func newMsg() string {
	return "bingo"
}

func newEcho() *echo.Echo {
	return echo.New()
}

func newWebApp(config *conf.Config, e *echo.Echo, blog *controller.BlogController) *WebApp {
	return &WebApp{
		e,
		config,
		blog,
	}
}

func loadTomlConf() (*conf.Config, error) {
	var conf conf.Config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
