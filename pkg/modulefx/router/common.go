package router

import (
	
	"net/http"


	"github.com/gin-gonic/gin"
	
)

const (
	CONTENT_JSON = "CONTENT_JSON"
	CONTENT_HTML = "CONTENT_HTML"
)


func responseWithSuccess[T any](gCtx *gin.Context, response T, contentType string) {
	switch contentType {
	case CONTENT_HTML:
		gCtx.AbortWithStatus(http.StatusOK)
	default:
		gCtx.Abort()
		gCtx.PureJSON(http.StatusOK, SuccessResponse[T]{Data: response})
	}
}


func responseWithError(gCtx *gin.Context, realsStatus int, err error) {
	
	gCtx.AbortWithStatusJSON(http.StatusOK, ErrorResponse{
		Error: ErrorDetails{
			Code:   realsStatus,
			Reason: err.Error(),
			Domain: "register_course_check",
		},
	})
}

