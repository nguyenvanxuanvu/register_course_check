package router

import (
	"context"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/controller"
	db_config "github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/dbconfig"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)


func NewHttpRouter(controller controller.Controller, dbConfig db_config.DBConfig) *gin.Engine {
	gin.SetMode(viper.GetString("debug.gin"))

	r := gin.Default()
	registerHandler(r, controller, dbConfig)
	return r
}

func registerHandler(r gin.IRoutes, controller controller.Controller, dbConfig db_config.DBConfig) {
	

	r.POST("/register_course/check",
		getControllerEndPoint("Check", getCheckRequestDTO, controller.Check, CONTENT_JSON))
	r.POST("/register_course/suggestion",
		getControllerEndPoint("Suggestion", getSuggestionRequestDTO, controller.Suggestion, CONTENT_JSON))
}


func getControllerEndPoint[Req any, Resp any](method string,
	getRequestFunc func(*gin.Context) (Req, error),
	controllerFunc func(context.Context, Req) (Resp, error),
	respContentType string,
) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		
		
		ctx := gCtx.Request.Context()
		request, err := getRequestFunc(gCtx)
		
		if err != nil {
			responseWithError(gCtx, http.StatusServiceUnavailable, err)
			return
		}
		response, err := controllerFunc(ctx, request)
		if err != nil {
			responseWithError(gCtx, http.StatusServiceUnavailable, err)
			return
		}
		
		responseWithSuccess(gCtx, response, respContentType)
	}
}





func getCheckRequestDTO(ctx *gin.Context) (*dto.CheckRequestDTO, error) {
	
	checkRequestDTO := dto.CheckRequestDTO{}
	
	if err := ctx.BindJSON(&checkRequestDTO); err != nil {
		return nil, err
	}
	
	return &checkRequestDTO, nil
}

func getSuggestionRequestDTO(ctx *gin.Context) (*dto.SuggestionRequestDTO, error) {
	
	suggestionRequestDTO := dto.SuggestionRequestDTO{}
	
	if err := ctx.BindJSON(&suggestionRequestDTO); err != nil {
		return nil, err
	}
	
	return &suggestionRequestDTO, nil
}
