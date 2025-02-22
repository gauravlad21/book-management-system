package controller

import (
	"context"
	"sync"

	"github.com/gauravlad21/book-management-system/dbhelper"
	"github.com/gauravlad21/book-management-system/external_resources/kafka"
	epredis "github.com/gauravlad21/book-management-system/external_resources/redis"
	"github.com/gauravlad21/book-management-system/service"
)

var cache epredis.CacheInterface
var dbOpsIf dbhelper.DbOperationsIF
var serviceRepo service.ServiceIF

func InitializeHandlers() {
	if cache == nil {
		cache = epredis.GetRedisClient()
	}

	if dbOpsIf == nil {
		dbOpsIf = dbhelper.GetDbOps()
	}

	if serviceRepo == nil {
		serviceRepo = service.New(dbOpsIf, cache)
	}
}

func StartupHook(ctx context.Context) {
	if serviceRepo == nil {
		InitializeHandlers()
	}

	kafka.InitKafkaProducer()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		kafka.StartConsumer()
	}()
}
