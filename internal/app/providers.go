package app

import (
	"log"
	api "yandex/second/package/api"
	db "yandex/second/package/database"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config *AppConfig
	Router *gin.Engine
}

func (a *App) Run() {
	logger := log.New(log.Writer(), "app: ", log.LstdFlags)
	logger.Println(("Starting server on " + a.Config.Http.Host + ":" + a.Config.Http.Port))
	a.Router.Run(a.Config.Http.Host + ":" + a.Config.Http.Port)
}

type HttpConfig struct {
	Host string
	Port string
}

type AppConfig struct {
	Http HttpConfig
}

func provideAppConfig() *AppConfig {
	return &AppConfig{
		Http: HttpConfig{
			Host: "localhost",
			Port: "8080",
		},
	}
}

func provideRouter() *gin.Engine {
	r := gin.Default()
	database, _ := db.ProvideDatabase()
	api.CarsHandlers(r, database)
	return r
}

func NewApp() *App {
	return &App{Config: provideAppConfig(), Router: provideRouter()}
}
