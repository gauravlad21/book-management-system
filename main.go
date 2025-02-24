package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gauravlad21/book-management-system/commonutility"
	urlmap "github.com/gauravlad21/book-management-system/urls_mappings"

	"github.com/gauravlad21/book-management-system/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/gauravlad21/book-management-system/docs" // Import the generated docs
)

func startServer(ctx context.Context, port string) {
	server := gin.New()
	server.Use(gin.Recovery())

	m := urlmap.GetUrlMaps()
	for _, urlMap := range m {
		url := urlMap.Url
		switch urlMap.Method {
		case urlmap.GET:
			server.GET(url, urlMap.Handler)
		case urlmap.POST:
			server.POST(url, urlMap.Handler)
		case urlmap.DELETE:
			server.DELETE(url, urlMap.Handler)
		case urlmap.PUT:
			server.PUT(url, urlMap.Handler)
		case urlmap.PATCH:
			server.PATCH(url, urlMap.Handler)
		}
	}

	server.Run(":" + port) // ":5002"
}

func initAndStartServer() {
	ctx := context.Background()
	commonutility.GetLogger()
	controller.InitializeHandlers()
	controller.StartupHook(ctx)
	port := viper.GetString("port")
	commonutility.GetLogger().Info(fmt.Sprintf("starting server on port: '%v'", port))
	startServer(ctx, port)
}

// @title Books Management API
// @version 1.0
// @description This is a Books Management API using Gin, GORM, Redis, and Kafka.
// @host 13.48.212.214
// @BasePath /
func main() {

	defaultPath := "default-path"
	var configPath string
	flag.StringVar(&configPath, "config", defaultPath, "local config path")

	flag.Parse()

	if configPath != defaultPath {
		commonutility.ReadConfigFile(configPath)
	}
	initAndStartServer()
}
