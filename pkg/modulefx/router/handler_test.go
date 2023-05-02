package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
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
	suite.handler = NewHttpRouter(controller, suite.dbConfig, authen)
}

func (suite *RouterTestSuite) TestUpdateCourseCondition() {

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/update_course_condition", nil)
	suite.handler.ServeHTTP(resp, req)

	//
	suite.Assert().Equal(http.StatusOK, resp.Code)

	jsonResp := resp.Body.String()
	errResp := &ErrorResponse{}
	json.Unmarshal([]byte(jsonResp), errResp)

	suite.Assert().Equal(http.StatusForbidden, errResp.Error.Code)
	suite.Assert().Equal("NOT_FOUND_API_KEY", errResp.Error.Reason)

	// 	//unauthenticated zpi response failed
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", nil)
	// 	req.Header.Set("Cookie", "zlp_token=unauthen;")

	// 	suite.sessionClient.EXPECT().
	// 		GetSession(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(nil, errors.New("session gone"))
	// 	suite.handler.ServeHTTP(resp, req)

	// 	suite.Assert().Equal(http.StatusOK, resp.Code)

	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusForbidden, errResp.Error.Code)
	// 	suite.Assert().Equal(common.REQUEST_IS_NOT_AUTHENTICATED, errResp.Error.Reason)

	// 	//validate request failed
	// 	redeemRequests := []dto.RedeemRequestDTO{
	// 		{
	// 			UserId:      1,
	// 			RefTransId:  "1",
	// 			AppId:       1,
	// 			GiftId:      1,
	// 			Amount:      0,
	// 			RequestTime: 1,
	// 		},
	// 		{
	// 			UserId:      1,
	// 			RefTransId:  "1",
	// 			AppId:       1,
	// 			GiftId:      0,
	// 			Amount:      1,
	// 			RequestTime: 1,
	// 		},
	// 		{
	// 			UserId:      1,
	// 			RefTransId:  "1",
	// 			AppId:       0,
	// 			GiftId:      1,
	// 			Amount:      1,
	// 			RequestTime: 1,
	// 		},
	// 	}

	// 	testCases := [][]string{
	// 		{"/cashier-gift-shop/v1/redeem", "INVALID_AMOUNT"},
	// 		{"/cashier-gift-shop/v1/redeem", "INVALID_GIFT_ID"},
	// 		{"/cashier-gift-shop/v1/redeem", "INVALID_APP_ID"},
	// 	}
	// 	for idx, testCase := range testCases {
	// 		resp = httptest.NewRecorder()
	// 		data, _ := json.Marshal(redeemRequests[idx])
	// 		reader := bytes.NewReader(data)
	// 		req, _ = http.NewRequest("POST", testCase[0], reader)
	// 		req.Header.Set("Cookie", "zlp_token=legit;")

	// 		suite.mockSessionSuccess()
	// 		suite.handler.ServeHTTP(resp, req)

	// 		//
	// 		suite.Assert().Equal(http.StatusOK, resp.Code)

	// 		jsonResp = resp.Body.String()
	// 		errResp = &ErrorResponse{}
	// 		json.Unmarshal([]byte(jsonResp), errResp)

	// 		suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 		suite.Assert().Equal(testCase[1], errResp.Error.Reason)
	// 	}

	// 	//get config fail

	// 	redeemRequest := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1",
	// 		AppId:       1,
	// 		GiftId:      10000,
	// 		Amount:      1,
	// 		RequestTime: 1,
	// 	}
	// 	data, _ := json.Marshal(redeemRequest)
	// 	reader := bytes.NewReader(data)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)

	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 	suite.Assert().Equal("NOT_FOUND_GIFT_CONFIG", errResp.Error.Reason)

	// 	// not appliaple gift

	// 	redeemRequest1 := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1",
	// 		AppId:       12,
	// 		GiftId:      2,
	// 		Amount:      50000,
	// 		RequestTime: 1,
	// 	}
	// 	data1, _ := json.Marshal(redeemRequest1)
	// 	reader1 := bytes.NewReader(data1)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader1)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)
	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 	suite.Assert().Equal("NOT_APPLIABLE_GIFT", errResp.Error.Reason)

	// 	// reward status is not in stock
	// 	redeemRequest2 := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1",
	// 		AppId:       12,
	// 		GiftId:      1,
	// 		Amount:      50000,
	// 		RequestTime: 1,
	// 	}
	// 	data2, _ := json.Marshal(redeemRequest2)
	// 	reader2 := bytes.NewReader(data2)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader2)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()
	// 	suite.mockRiskNormal()
	// 	suite.mockGetInforOfRewardsSuccess()

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)

	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 	suite.Assert().Equal("INACTIVE_REWARD", errResp.Error.Reason)

	// 	//get infor of reward failed

	// 	redeemRequest3 := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1",
	// 		AppId:       12,
	// 		GiftId:      1,
	// 		Amount:      50000,
	// 		RequestTime: 1,
	// 	}
	// 	data3, _ := json.Marshal(redeemRequest3)
	// 	reader3 := bytes.NewReader(data3)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader3)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()
	// 	suite.mockRiskNormal()

	// 	suite.rewardStoreClient.EXPECT().
	// 		GetRewardDetail(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(nil, errors.New("store gone"))

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)

	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 	suite.Assert().Equal("store gone", errResp.Error.Reason)

	// 	// rule fail
	// 	viper.Set("rule-engine.client-id", 1)
	// 	redeemRequest4 := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1xyz",
	// 		AppId:       12,
	// 		GiftId:      8,
	// 		Amount:      50000,
	// 		RequestTime: 1,
	// 	}
	// 	data4, _ := json.Marshal(redeemRequest4)
	// 	reader4 := bytes.NewReader(data4)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader4)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()
	// 	suite.mockRiskNormal()
	// 	suite.mockGetInforOfRewardsSuccess1()

	// 	suite.ruleEngineClient.EXPECT().
	// 		CheckAndUpdate(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(nil, errors.New("error"))
	// 	suite.kafkaProducer.EXPECT().ProduceRedeemLog(gomock.Any(), gomock.Any()).Times(1)

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)
	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 	suite.Assert().Equal("FAIL_RULE", errResp.Error.Reason)

	// 	// redeem point fail
	// 	redeemRequest5 := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1",
	// 		AppId:       12,
	// 		GiftId:      1,
	// 		Amount:      50000,
	// 		RequestTime: 1,
	// 	}
	// 	data5, _ := json.Marshal(redeemRequest5)
	// 	reader5 := bytes.NewReader(data5)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader5)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()
	// 	suite.mockRiskNormal()
	// 	suite.mockGetInforOfRewardsSuccess1()

	// 	suite.pointEngineClient.EXPECT().
	// 		RedeemPoint(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(nil, errors.New("redeem point fail"))
	// 	suite.kafkaProducer.EXPECT().ProduceRedeemLog(gomock.Any(), gomock.Any()).Times(1)

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)
	// 	jsonResp = resp.Body.String()
	// 	errResp = &ErrorResponse{}
	// 	json.Unmarshal([]byte(jsonResp), errResp)

	// 	suite.Assert().Equal(http.StatusServiceUnavailable, errResp.Error.Code)
	// 	suite.Assert().Equal("redeem point fail", errResp.Error.Reason)

	// 	// successfully

	// 	redeemRequest6 := dto.RedeemRequestDTO{

	// 		UserId:      1,
	// 		RefTransId:  "1",
	// 		AppId:       12,
	// 		GiftId:      1,
	// 		Amount:      50000,
	// 		RequestTime: 1,
	// 	}
	// 	data6, _ := json.Marshal(redeemRequest6)
	// 	reader6 := bytes.NewReader(data6)
	// 	resp = httptest.NewRecorder()
	// 	req, _ = http.NewRequest("POST", "/cashier-gift-shop/v1/redeem", reader6)
	// 	req.Header.Set("Cookie", "zlp_token=legit;")

	// 	suite.mockSessionSuccess()
	// 	suite.mockRiskNormal()
	// 	suite.mockGetInforOfRewardsSuccess1()
	// 	suite.mockRedeemPointSuccess()

	// 	suite.kafkaProducer.EXPECT().ProduceRedeemLog(gomock.Any(), gomock.Any()).Times(1)

	// 	suite.handler.ServeHTTP(resp, req)

	// 	//
	// 	suite.Assert().Equal(http.StatusOK, resp.Code)
	// 	jsonResp = resp.Body.String()b
	// 	dataResp := &SuccessResponse[dto.RedeemResponseDTO]{}
	// 	json.Unmarshal([]byte(jsonResp), dataResp)

	// 	suite.Assert().Equal(dataResp.Data.ReceivedRewards[0].PromotionType, int32(1))
	// 	suite.Assert().Equal(dataResp.Data.CurrentPoint, uint64(10))

	// }

	// func (suite *RouterTestSuite) mockSessionSuccess() {
	// 	suite.sessionClient.EXPECT().
	// 		GetSession(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(&zpi_session_grpc.Session{
	// 			ZalopayId: "123",
	// 		}, nil)
	// }

	// func (suite *RouterTestSuite) mockRiskNormal() {
	// 	suite.riskClient.EXPECT().
	// 		SendRequestV2(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(&zalopay.GrpcBaseResponse{
	// 			Response: jsonMustMarshal(risk.RiskResponse{
	// 				Data: risk.ResponseData{
	// 					Code: 1,
	// 				},
	// 			}),
	// 		}, nil)

	// }

	// func (suite *RouterTestSuite) mockRiskCasualAbuser() {
	// 	suite.riskClient.EXPECT().
	// 		SendRequestV2(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(&zalopay.GrpcBaseResponse{
	// 			Response: jsonMustMarshal(risk.RiskResponse{
	// 				Data: risk.ResponseData{
	// 					Code: 0,
	// 				},
	// 			}),
	// 		}, nil)
	// }

	// func (suite *RouterTestSuite) mockGetRewardsSuccess() {
	// 	suite.rewardStoreClient.EXPECT().
	// 		GetRewardSection(gomock.Any(), gomock.Any(), gomock.Any()).
	// 		Return(&reward_store_grpc.SearchRewardResponse{
	// 			Rewards: []*reward_store_grpc.RewardDisplayDTO{
	// 				{
	// 					RewardStoreId: 1,
	// 					DisplayTitle:  "Giam 10K",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 2,
	// 					DisplayTitle:  "Giam 20K",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 3,
	// 					DisplayTitle:  "Giam 30K",
	// 					BrandName:     "TTDV",
	// 				},

	// 				{
	// 					RewardStoreId: 4,
	// 					DisplayTitle:  "Giam 40K",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 5,
	// 					DisplayTitle:  "Giam 50K",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 6,
	// 					DisplayTitle:  "Giam 60K",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 7,
	// 					DisplayTitle:  "Giam 100K",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 8,
	// 					DisplayTitle:  "Giam 10K - Rule",
	// 					BrandName:     "TTDV",
	// 				},
	// 				{
	// 					RewardStoreId: 69,
	// 					DisplayTitle:  "Giam 69K - not existed",
	// 					BrandName:     "TTDV",
	// 				},
	// 			},
	// 		}, nil)
	// }
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, &RouterTestSuite{})
}
