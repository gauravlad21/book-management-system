package commonutility

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) context.Context {
	// should contains reqeust specific values in context
	return context.Background()
}
