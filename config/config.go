package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const envPerfix string = "GIN_"

/*Init : 初始化配置*/
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no such .env file, use system env")
	} else {
		log.Println("use .env file")
	}
}

// AppConfig ...
type AppConfig struct {
	Name      string `json:"name"`
	RunMode   string `json:"run_mode"`
	Addr      string `json:"addr"`
	JwtSecret string `json:"jwt_secret"`
	// JWTExpirationTime day
	JwtExpirationTime int `json:"jwt_expiration_time"`
}

// MySQLConfig ...
type MySQLConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	Charset      string
	ShowSQL      bool
	MaxIdleConns int
	MaxOpenConns int
}

// RedisConfig ...
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// CacheConfig ...
type CacheConfig struct {
	Driver string `json:"driver"`
	Prefix string `json:"prefix"`
}

// SentinelRuleConfig ...
type SentinelRuleConfig struct {
	Resource        string `json:"resource"`
	MetricType      string `json:"metric_type"`
	ControlBehavior string `json:"control_behavior"`
	Count           int64  `json:"count"`
}

// Config global config
// include common and biz config
type Config struct {
	// common
	App           AppConfig            `json:"app"`
	MySQL         MySQLConfig          `json:"mysql"`
	Redis         RedisConfig          `json:"redis"`
	Cache         CacheConfig          `json:"cache"`
	SentinelRules []SentinelRuleConfig `json:"sentinel_rules"`
}

/*Database : 数据库配置内容*/
func Database() MySQLConfig {
	Host := Getenv("DB_HOST")
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := Getenv("DB_PORT")
	if Port == "" {
		Port = "3306"
	}
	User := Getenv("DB_USER")
	Password := Getenv("DB_PASSWORD")
	Name := Getenv("DB_NAME")
	Charset := Getenv("DB_CHARSET")
	if Charset == "" {
		Charset = "utf8mb4"
	}
	ShowSQL := Getenv("DB_SHOW_SQL")
	MaxIdleConns, err := strconv.Atoi(Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		MaxIdleConns = 2
	}
	MaxOpenConns, err := strconv.Atoi(Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		MaxOpenConns = 10
	}
	return MySQLConfig{
		Host:         Host,
		Port:         Port,
		User:         User,
		Password:     Password,
		Name:         Name,
		Charset:      Charset,
		ShowSQL:      ShowSQL == "true",
		MaxIdleConns: MaxIdleConns,
		MaxOpenConns: MaxOpenConns,
	}
}

/*Redis : Redis 配置内容*/
func Redis() RedisConfig {
	Host := Getenv("REDIS_HOST")
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := Getenv("REDIS_PORT")
	if Port == "" {
		Port = "6379"
	}
	Password := Getenv("REDIS_PASSWORD")
	DB, err := strconv.Atoi(Getenv("REDIS_DB"))
	if err != nil {
		DB = 0
	}
	return RedisConfig{
		Host:     Host,
		Port:     Port,
		Password: Password,
		DB:       DB,
	}
}

/*Getenv : 获取环境变量*/
func Getenv(name string) string {
	return os.Getenv(envPerfix + name)
}
