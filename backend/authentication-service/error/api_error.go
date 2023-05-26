package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBadRequest(ctx *gin.Context, err error, errorMessage string, args ...interface{}) {
	apiError := APILayerError.WrapWithUserMessage(err, errorMessage, args)
	ctx.JSON(http.StatusBadRequest, apiError.JSON())
}

func HandleBadRequestWithStatusCode(ctx *gin.Context, statusCode int, err error, errorMessage string, args ...interface{}) {
	apiError := APILayerError.WrapWithUserMessage(err, errorMessage, args)
	ctx.JSON(statusCode, apiError.JSON())
}
