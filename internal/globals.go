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

	MSG_DB_FAILED      = "Internal server error. Something went wrong on our end. Please try again later."
	MSG_WRONG_PARAMS   = "Invalid input. Please check your data and try again."
	MSG_RATE_LIMIT     = "Bạn truy cập quá nhanh."
	MSG_TOKEN_REQUIRED = "Token không tồn tại"
	MSG_TOKEN_EXPIRED  = "Token hết hạn"
	MSG_INVALID_TOKEN  = "Token không hợp lệ"
	MSG_SYSTEM_BUSY    = "Hệ thống đang bận, vui lòng thử lại sau ít phút"
	MSG_SYSTEM_ERROR   = "System error. An unexpected error occurred. Please try again later."
)

var Envs = InitEnvVars()
var SysStatus = InitSystemStatus()
var Log = NewLogger()

var Db = NewSQLDB()
var Pool = NewThreadPool()
var Validator = validator.New()
var Keys = InitKeys(Envs.IsProduction)
var MailEnvs = InitMailSender()

type AppKeys struct {
	ACCESS_TOKEN_SECRET        string
	REFRESH_TOKEN_SECRET       string
	INSIDE_ACCESS_TOKEN_SECRET string
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
type MailConfig struct {
	SMTPHost     string
	SMTPPort     string
	MaiFrom      string
	MailPassword string
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

func InitMailSender() *MailConfig {
	return &MailConfig{
		MaiFrom:      "nguyenbac872001@gmail.com",
		SMTPHost:     "smtp.gmail.com",
		SMTPPort:     "587",
		MailPassword: "fuetfgukqnvqdgtt",
	}
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
		ACCESS_TOKEN_SECRET:        "2250bcba-9227-418d-9d23-6b5492d81de1",
		REFRESH_TOKEN_SECRET:       "8b93f810-4c3a-452d-9a1a-75101ba7b080",
		INSIDE_ACCESS_TOKEN_SECRET: "DnEMRqGdoR07HbY-l9xRXWuCo3dnYzNzqoAoOqEZO_EpAOKMql",
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

// "Access denied. You don't have permission to perform this action."
