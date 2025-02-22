package controller

import (
	"context"

	"github.com/gauravlad21/book-management-system/dbhelper"
	"github.com/gauravlad21/book-management-system/service"
)

var dbOpsIf dbhelper.DbOperationsIF
var serviceRepo service.ServiceIF

func InitializeHandlers() {
	if dbOpsIf == nil {
		dbOpsIf = dbhelper.GetDbOps()
	}

	if serviceRepo == nil {
		serviceRepo = service.New(dbOpsIf)
	}
}

func StartupHook(ctx context.Context) {
	if serviceRepo == nil {
		InitializeHandlers()
	}
}
