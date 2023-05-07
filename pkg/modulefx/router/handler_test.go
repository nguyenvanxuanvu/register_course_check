package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/authen"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/controller"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/dbconfig"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/repository"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/service"
	"github.com/nguyenvanxuanvu/register_course_check/testing/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
	db       *sqlx.DB
	handler  *gin.Engine
	mockCtrl *gomock.Controller
	dbConfig dbconfig.DBConfig
	client   *mocks.MockClient
	cache    *mocks.MockCacheService
}
func (suite *RouterTestSuite) SetupSuite() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		suite.T().Fatal(err.Error())
	}
	//insert data
	stmts, err := os.ReadFile("./test/script.sql")
	if err != nil {
		suite.T().Fatal(err.Error())
		return
	}
	_, err = db.Exec(string(stmts))
	if err != nil {
		suite.T().Fatal(err.Error())
		return
	}

	////
	suite.db = db

	configRepo := repository.NewConfigRepository(db)
	repo := repository.NewRepository(db)
	suite.dbConfig, err = dbconfig.NewDBConfig(configRepo)
	if err != nil {
		suite.T().Fatal(err.Error())
	}

	
	suite.cache = mocks.NewMockCacheService(suite.mockCtrl)
	
	//client
	suite.client = mocks.NewMockClient(suite.mockCtrl)
	client := client.NewClient()
	authen := authen.NewAuthenticator()
	service := service.NewRegisterCourseCheckService(suite.dbConfig, repo, client, suite.cache)
	controller := controller.NewController(service)

	viper.Set("debug.gin", "debug")
	viper.Set("authen.api-key", "abcd")
	suite.handler = NewHttpRouter(controller, suite.dbConfig, authen)
}

func (suite *RouterTestSuite) TestUpdateCourseCondition() {
	// not found api key
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/update_course_condition", nil)
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp := resp.Body.String()
	errResp := &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	suite.Assert().Equal("NOT_FOUND_API_KEY", errResp.Error.Reason)


	// Wrong api key

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/update_course_condition", nil)
	req.Header.Add("Apikey", "abc")
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp = resp.Body.String()
	errResp = &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	suite.Assert().Equal(http.StatusForbidden, errResp.Error.Code)
	suite.Assert().Equal("WRONG_API_KEY", errResp.Error.Reason)

	// update successful
	updateCourseConditionRequest := []dto.CourseConditionConfig{
		{
			CourseId: "test",
			Condition: &dto.CourseCondition{
				Op: "or",
				Leaves: nil,
				Course: &dto.CourseConditionInfo{
					CourseDesId: "test1",
					Type: 1,
				},
			},
		},
	}
	data, _ := json.Marshal(updateCourseConditionRequest)
	reader := bytes.NewReader(data)
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/update_course_condition", reader)
	req.Header.Add("Apikey", "abcd")
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp = resp.Body.String()
	errResp = &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	dataResp := &SuccessResponse[bool]{}
	json.Unmarshal([]byte(jsonResp), dataResp)
	suite.Assert().Equal(dataResp.Data, true)

}






func (suite *RouterTestSuite) TestSuggestion() {
	// not found api key
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register_course/suggestion", nil)
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp := resp.Body.String()
	errResp := &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	suite.Assert().Equal("NOT_FOUND_API_KEY", errResp.Error.Reason)


	// Wrong api key

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/register_course/suggestion", nil)
	req.Header.Add("Apikey", "abc")
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp = resp.Body.String()
	errResp = &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	suite.Assert().Equal(http.StatusForbidden, errResp.Error.Code)
	suite.Assert().Equal("WRONG_API_KEY", errResp.Error.Reason)

	// update successful
	suggestionRequest := dto.SuggestionRequestDTO{
		StudentId: "1915982",
		Semester: 192,
	}
	data, _ := json.Marshal(suggestionRequest)
	reader := bytes.NewReader(data)
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/register_course/suggestion", reader)
	req.Header.Add("Apikey", "abcd")
	
	suite.cache.EXPECT().GetStudentInfo(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp = resp.Body.String()
	errResp = &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	dataResp := &SuccessResponse[dto.SuggestionResponseDTO]{}
	json.Unmarshal([]byte(jsonResp), dataResp)
	suite.Assert().Equal(dataResp.Data.MinCredit, 4)

}



func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, &RouterTestSuite{})
}
