package internal

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"github.com/shettyh/threadpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	COST_APPOINTMENT = "APPOINTMENT_TIME_DIFF"
	COST_DISTANCE    = "DISTANCE_DIFF_EMP_TASK"
	COST_MATCH_SKILL = "MATCH_SKILLS"
)

const (
	SERVICE_NAME = "hi-ecom-fconnect-v2-api"
)
const (
	INF_COST            = 1e9
	NEG_INF_COST        = -1e9
	CODE_DB_FAILED      = 300
	CODE_WRONG_PARAMS   = 400
	CODE_RATE_LIMIT     = 401
	CODE_TOKEN_REQUIRED = 1003
	CODE_TOKEN_EXPIRED  = 1001
	CODE_INVALID_TOKEN  = 1002
	CODE_SYSTEM_BUSY    = 300
	CODE_SYSTEM_ERROR   = 301

	MSG_DB_FAILED      = "Kết nối hệ thống lỗi, vui lòng thử lại sau ít phút"
	MSG_WRONG_PARAMS   = "Sai thông tin đầu vào, vui lòng kiểm tra lại thông tin"
	MSG_RATE_LIMIT     = "Bạn truy cập quá nhanh."
	MSG_TOKEN_REQUIRED = "Token không tồn tại"
	MSG_TOKEN_EXPIRED  = "Token hết hạn"
	MSG_INVALID_TOKEN  = "Token không hợp lệ"
	MSG_SYSTEM_BUSY    = "Hệ thống đang bận, vui lòng thử lại sau ít phút"
	MSG_SYSTEM_ERROR   = "Có lỗi trong quá trình xử lý, vui lòng thử lại sau ít phút"
)

var Envs = InitEnvVars()
var SysStatus = InitSystemStatus()
var Log = NewLogger()

var Db = NewSQLDB()
var Pool = NewThreadPool()
var Validator = validator.New()
var Domains = InitAPIDomains(Envs.IsProduction)
var Eps = InitAPIEndpoints()
var Keys = InitKeys(Envs.IsProduction)
var Brokers = NewBrokers(Envs.IsProduction)
var KafkaTopicName = NewKafkaTopicName(Envs.IsProduction)
var KafkaTopicNameAll = NewKafkaTopicNameAll(Envs.IsProduction)
var OrderState = InitOrderState()

type OrderStateStruct struct {
	AccepTask           string
	AssignTask          string
	AssignTenant        string
	Cancel              string
	CheckIn             string
	CheckOutFail        string
	CheckOutSuccess     string
	ConfirmOrderPrice   string
	CustomerCancelOrder string
	Monitoring          string
	OutCase             string
	Recall              string
	DictStateDes        map[string]string
	TaskStateFinal      []string
}

func InitOrderState() *OrderStateStruct {
	// dict state des
	dictStateDes := make(map[string]string)
	dictStateDes["ACCEPT_TASK"] = "Nhận tuyến"
	dictStateDes["ASSIGNED_TASK"] = "Phân tuyến"
	dictStateDes["ASSIGNED_TENANT"] = "Hệ thống đã phân công tới đối tác"
	dictStateDes["CANCEL"] = "Hẹn lại khách hàng"
	dictStateDes["CHECK_IN"] = "Check in nhà khách hàng"
	dictStateDes["CHECK_OUT_FAIL"] = "Đóng ca - phục vụ thất bại"
	dictStateDes["CHECK_OUT_SUCCESS"] = "Đóng ca - phục vụ thành công"
	dictStateDes["CONFIRM_ORDER_PRICE"] = "Nhân viên xác nhận giá đơn hàng"
	dictStateDes["CUSTOMER_CANCEL_ORDER"] = "Khách hàng hủy đơn hàng"
	dictStateDes["MORNITORING"] = "Đã xử lý đang theo dõi"
	dictStateDes["OUT_CASE"] = "Nhả tuyến"
	dictStateDes["RECALL"] = "Nhả tuyến"
	// task state final
	taskStateFinal := []string{"CUSTOMER_CANCEL_ORDER", "CANCEL", "CHECK_OUT_FAIL", "CHECK_OUT_SUCCESS"}

	orderState := &OrderStateStruct{
		AccepTask:           "ACCEPT_TASK",
		AssignTask:          "ASSIGNED_TASK",
		AssignTenant:        "ASSIGN_TENANT",
		Cancel:              "CANCEL",
		CheckIn:             "CHECK_IN",
		CheckOutFail:        "CHECK_OUT_FAIL",
		CheckOutSuccess:     "CHECK_OUT_SUCCESS",
		ConfirmOrderPrice:   "CONFIRM_ORDER_PRICE",
		CustomerCancelOrder: "CUSTOMER_CANCEL_ORDER",
		Monitoring:          "MONITORING",
		OutCase:             "OUT_CASE",
		Recall:              "RECALL",
		DictStateDes:        dictStateDes,
		TaskStateFinal:      taskStateFinal,
	}
	return orderState
}

type AppKeys struct {
	TOKEN_SECRET_KEY_WEB     string
	OrderClientKey           string
	OrderServiceKey          string
	PORTAL_FCONNECT_AUTH_KEY string
	LocalServicesKey         map[string]string
	PartnerServiceClientKey  string
	FcOrderServiceClientKey  string
	ServiceKey               string
	AddressClientKey         string
	AddressSecretKey         string
	HiFPTEcomClientKey       string
	HiFPTEcomSecrectKey      string
}
type EnvVars struct {
	SqlHost      string `mapstructure:"DB_HOST"`
	SqlUser      string `mapstructure:"DB_USERNAME"`
	SqlPassword  string `mapstructure:"DB_PASSWORD"`
	SqlDBName    string `mapstructure:"DB_DATABASE"`
	SqlPort      int    `mapstructure:"DB_PORT"`
	ServicePort  string `mapstructure:"SERVICE_PORT"`
	IsProduction bool   `mapstructure:"USE_PRODUCTION"`
}
type ApiDomains struct {
	PublicHiFPT         string
	PrivateHiFPT        string
	UploadUrlHiFPT      string
	UrlStatic           string
	UrlHiFConnectOrder  string
	UrlCustomerProvider string
	UrlHiEcomPromotion  string
	UrlHiEcomBilling    string
}
type GroupEndpointHiFPTPromotion struct {
	GetListSkuFrPromotionCode string
}

type GroupEndpointHiFPTBilling struct {
	GetListCsatCSScoreFrOrderId string
}

type ApiEndpoints struct {
	GetTaskEp      string
	GetTenantEp    string
	AssignTaskEp   string
	TaskInfo       string
	GetProvince    string
	GetDistrict    string
	GetWard        string
	GetStreet      string
	GetReport      string
	HiFPTPromotion GroupEndpointHiFPTPromotion
	HiFPTBilling   GroupEndpointHiFPTBilling
}
type SystemStatus struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}
type AllSystemStatus struct {
	DbFailed      *SystemStatus
	WrongParams   *SystemStatus
	RateLimit     *SystemStatus
	TokenRequired *SystemStatus
	TokenExpired  *SystemStatus
	InvalidToken  *SystemStatus
	SystemBusy    *SystemStatus
	SystemError   *SystemStatus
}

func InitSystemStatus() *AllSystemStatus {
	return &AllSystemStatus{
		DbFailed: &SystemStatus{
			Status: CODE_DB_FAILED,
			Msg:    MSG_DB_FAILED,
		},
		WrongParams: &SystemStatus{
			Status: CODE_WRONG_PARAMS,
			Msg:    MSG_WRONG_PARAMS,
		}, RateLimit: &SystemStatus{
			Status: CODE_RATE_LIMIT,
			Msg:    MSG_RATE_LIMIT,
		}, TokenRequired: &SystemStatus{
			Status: CODE_TOKEN_REQUIRED,
			Msg:    MSG_TOKEN_REQUIRED,
		}, TokenExpired: &SystemStatus{
			Status: CODE_TOKEN_EXPIRED,
			Msg:    MSG_TOKEN_EXPIRED,
		}, InvalidToken: &SystemStatus{
			Status: CODE_INVALID_TOKEN,
			Msg:    MSG_INVALID_TOKEN,
		}, SystemBusy: &SystemStatus{
			Status: CODE_SYSTEM_BUSY,
			Msg:    MSG_SYSTEM_BUSY,
		}, SystemError: &SystemStatus{
			Status: CODE_SYSTEM_ERROR,
			Msg:    MSG_SYSTEM_ERROR,
		},
	}
}
func InitKeys(isProduction bool) *AppKeys {
	keys := &AppKeys{
		TOKEN_SECRET_KEY_WEB:     "FCONNECT2_hehelilifsdfjjfjglilieiruwoeiurMsIjH",
		OrderClientKey:           "49e7828e-adf1-47fe-a355-553332e9da1f",
		OrderServiceKey:          "5850405d-0242-49e1-808a-c2c4bbbfae84",
		PORTAL_FCONNECT_AUTH_KEY: "PORTAL_ldsjdhsajdhwclieiruwoeiurMsIjHdasdasd",
		PartnerServiceClientKey:  "074a8c0b-0478-49cf-897f-36493003451a",
		AddressClientKey:         "WebkitHiFPT",
		AddressSecretKey:         "48a970b5c2678204b590b6a997444b19",
		ServiceKey:               "2d51c978-6a34-4d33-a7b6-1ed7956e2583",
		HiFPTEcomClientKey:       "hifpt_ecom",
		HiFPTEcomSecrectKey:      "xxxxxxecom2021",
	}
	if isProduction {
		keys.PartnerServiceClientKey = "5c273964-3581-477c-b442-bc5cd32a3517"
		keys.ServiceKey = "b7ef02a7-ab9f-449f-9e4f-0b8ffacf3a95"

		keys.OrderClientKey = "fb76a30c-09fc-4889-ab47-dbbb23f50aec"
		keys.OrderServiceKey = "d7560895-9764-4209-90a4-ba7196180cf8"
	}
	keys.LocalServicesKey = map[string]string{
		"hi-fconnect-partner-management-api": keys.PartnerServiceClientKey,
	}
	return keys
}
func NewLogger() *zap.Logger {
	// Log vào stdOut
	writeSyncer := zapcore.AddSync(os.Stdout)
	// Thiết lập log
	loggerCore := zapcore.NewCore(logEndcoder(), writeSyncer, zap.DebugLevel)
	return zap.New(loggerCore, zap.AddCaller())
}
func logEndcoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
func InitAPIDomains(isProduction bool) *ApiDomains {
	// staging
	domain := &ApiDomains{
		PrivateHiFPT: "http://hifpt-api-stag.fpt.vn",
		PublicHiFPT:  "https://hifpt-api-stag.fpt.vn",
		// PublicHiFPT: "http://localhost:3003",
		UploadUrlHiFPT:      "img_stg",
		UrlStatic:           "https://hi-static.fpt.vn",
		UrlHiFConnectOrder:  "http://hi-fconnect-order-management-api-staging:80",
		UrlCustomerProvider: "http://hi-customer-provider-api-staging",
		UrlHiEcomPromotion:  "http://hi-ecom-promotion-api-staging:9003",
		UrlHiEcomBilling:    "http://hi-ecom-billing-api-staging:9004",
	}
	// production
	if isProduction {
		domain.PrivateHiFPT = "http://hifpt-api.fpt.vn"
		domain.PublicHiFPT = "https://hifpt-api.fpt.vn"
		domain.UploadUrlHiFPT = "img"
		domain.UrlHiFConnectOrder = "http://hi-fconnect-order-management-api:80"
		domain.UrlCustomerProvider = "http://hi-customer-provider-api"
		domain.UrlHiEcomPromotion = "http://hi-ecom-promotion-api:9003"
		domain.UrlHiEcomBilling = "http://hi-ecom-billing-api:9004"
	}
	return domain
}
func InitAPIEndpoints() *ApiEndpoints {
	prefixOrder := "/hi-fconnect-order-management-api"
	prefixPartner := "/hi-fconnect-partner-management-api"
	prefixAddressProvider := "/customer-provider/location"
	endpoints := &ApiEndpoints{
		GetTaskEp:    prefixOrder + "/orders/tasks-state",
		TaskInfo:     prefixOrder + "/orders/task-info",
		GetTenantEp:  prefixPartner + "/v1/local/tenant/get",
		AssignTaskEp: prefixPartner + "/v1/local/tenant/assign-task",
		GetProvince:  prefixAddressProvider + "/provinces",
		GetDistrict:  prefixAddressProvider + "/districts",
		GetWard:      prefixAddressProvider + "/wards",
		GetStreet:    prefixAddressProvider + "/streets",
		GetReport:    prefixOrder + "/reports",
		HiFPTPromotion: GroupEndpointHiFPTPromotion{
			GetListSkuFrPromotionCode: "/hi-ecom-promotion-api/v1/local/get-skus-fr-promotion-code",
		},
		HiFPTBilling: GroupEndpointHiFPTBilling{
			GetListCsatCSScoreFrOrderId: "/hi-ecom-billing-api/v1/local/recare-get-csat-fr-order-id",
		},
	}
	return endpoints
}
func InitEnvVars() *EnvVars {
	fmt.Println("LOADING ENVS...")
	envs := &EnvVars{}
	viper.SetConfigFile(".env")
	errEnvFile := viper.ReadInConfig()
	if errEnvFile != nil {
		viper.AutomaticEnv()
		viper.BindEnv("DB_HOST")
		viper.BindEnv("DB_USERNAME")
		viper.BindEnv("DB_PASSWORD")
		viper.BindEnv("DB_DATABASE")
		viper.BindEnv("DB_PORT")
		viper.BindEnv("USE_PRODUCTION")
		viper.BindEnv("SERVICE_KEY")
		viper.BindEnv("SERVICE_PORT")
	}
	if err := viper.Unmarshal(envs); err != nil {
		fmt.Println("Error viper.Unmarshal", err)
		fmt.Println("LOADING ENVS FAILED")
	}
	if envs.ServicePort == "" {
		envs.ServicePort = "8080"
	}
	fmt.Println("LOADING ENVS SUCCESS")
	return envs
}
func NewSQLDB() *gorm.DB {
	fmt.Println("LOADING MYSQL DB ...")
	DBDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&interpolateParams=true",
		Envs.SqlUser, Envs.SqlPassword, Envs.SqlHost, Envs.SqlPort, Envs.SqlDBName)
	DBDSN = DBDSN + "&loc=Asia%2FHo_Chi_Minh"
	DB, err := sql.Open("mysql", DBDSN)
	if err != nil {
		// cm.Log.Debug("Open mysql connection failed", zap.Error(err))
		fmt.Println("NewSQLDB", err)
		return nil
	}
	DB.SetConnMaxLifetime(time.Minute * 10)
	DB.SetMaxOpenConns(1000)
	DB.SetMaxIdleConns(1000)
	errPing := DB.Ping()
	if errPing != nil {
		fmt.Println("DB PING ERR: ", errPing)
		return nil
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}))
	if err != nil {
		fmt.Println("gorm DB connect failed", err.Error())
		return nil
	}
	fmt.Println("LOADING MYSQL DB SUCCESS...")
	return gormDB
}

func NewThreadPool() *threadpool.ThreadPool {
	fmt.Println("LOADING THREAD POOL ...")
	threadPool := threadpool.NewThreadPool(100, 3000)
	if threadPool == nil {
		fmt.Println("Failed")
	}
	fmt.Println("LOADING THREAD POOL SUCCESS...")
	return threadPool
}
func NewBrokers(useProduction bool) []string {
	if useProduction {
		return []string{"isc-kafka01:9092", "isc-kafka02:9092", "isc-kafka03:9092"}
	}
	return []string{"isc-kafka01:9092", "isc-kafka02:9092", "isc-kafka03:9092"}
}
func NewKafkaTopicName(useProduction bool) string {
	if useProduction {
		return "hifpt-hi-ecom-logs"
	}
	return "stag-hifpt-hi-ecom-logs"
}
func NewKafkaTopicNameAll(useProduction bool) string {
	if useProduction {
		return "hifpt-all-in-one-kafka"
	}
	return "stag-hifpt-all-in-one-kafka"
}
