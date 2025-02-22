package commonutility

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) context.Context {
	// should contains reqeust specific values in context
	return context.Background()
}

func GetCacheKey(id string) string {
	return "BookId::" + id
}

func GetAllBooksKey(limit, offset int) string {
	return fmt.Sprintf("AllBooks::limit:%v::offset:%v", limit, offset)
}

func GetAllBooksKeyPrefix() string {
	return "AllBooks*"
}
