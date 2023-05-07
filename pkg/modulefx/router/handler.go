package router

import (
	"context"
	"errors"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/authen"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/controller"
	db_config "github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/dbconfig"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/spf13/viper"
)

func NewHttpRouter(controller controller.Controller, dbConfig db_config.DBConfig, authen authen.Authenticator) *gin.Engine {
	gin.SetMode(viper.GetString("debug.gin"))

	r := gin.Default()
	registerHandler(r, controller, dbConfig, authen)
	return r
}

func registerHandler(r gin.IRoutes, controller controller.Controller, dbConfig db_config.DBConfig, authen authen.Authenticator) {

	r.POST("/register_course/check", authenticateMiddleware(authen),
		getControllerEndPoint("Check", getCheckRequestDTO, controller.Check, CONTENT_JSON))
	r.POST("/register_course/suggestion", authenticateMiddleware(authen),
		getControllerEndPoint("Suggestion", getSuggestionRequestDTO, controller.Suggestion, CONTENT_JSON))
	r.POST("/update_course_condition", authenticateMiddleware(authen),
		getControllerEndPoint("UpdateCourseCondition", getUpdateCourseConditionRequestDTO, controller.UpdateCourseCondition, CONTENT_JSON))
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
			responseWithError(gCtx, http.StatusServiceUnavailable, err, "")
			return
		}
		response, err := controllerFunc(ctx, request)
		if err != nil {
			responseWithError(gCtx, http.StatusServiceUnavailable, err, ErrToDescription(err))
			return
		}

		responseWithSuccess(gCtx, response, respContentType)
	}
}

func authenticateMiddleware(authenticator authen.Authenticator) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		apiKey := gCtx.Request.Header["Apikey"]
		if len(apiKey) == 0 {
			responseWithError(gCtx, http.StatusForbidden, errors.New(common.NOT_FOUND_API_KEY), ErrToDescription(errors.New(common.NOT_FOUND_API_KEY)))
			return
		}
		check := authenticator.Authen(apiKey[0])
		if check == false {
			responseWithError(gCtx, http.StatusForbidden, errors.New(common.WRONG_API_KEY), ErrToDescription(errors.New(common.WRONG_API_KEY)))
			return
		}
		gCtx.Next()
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

func getUpdateCourseConditionRequestDTO(ctx *gin.Context) ([]dto.CourseConditionConfig, error) {

	updateCourseConditionRequestDTO := []dto.CourseConditionConfig{}

	if err := ctx.BindJSON(&updateCourseConditionRequestDTO); err != nil {
		return nil, err
	}

	return updateCourseConditionRequestDTO, nil
}
