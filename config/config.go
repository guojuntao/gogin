package config

import (
	"os"
)

// TODO: 考虑使用 viper 管理配置
type Config struct {
	HttpPort string

	MongoDomain string
	MongoPort   string
	MongoAuth   string

	DbName         string
	ItemCollection string

	LogFile  string
	DebugLog bool
	TraceLog bool
}

var config Config

func init() {
	config = getConfig()
}

func GetConfig() Config {
	return config
}

func getConfig() (cfg Config) {
	cfg.HttpPort = getEnv("httpPort", "9999")
	cfg.MongoDomain = getEnv("mongoDomain", "localhost")
	cfg.MongoPort = getEnv("mongoPort", "27017")
	cfg.DbName = getEnv("db", "go-gin")
	cfg.ItemCollection = getEnv("itemCol", "item")

	cfg.LogFile = getEnv("logFile", "")
	cfg.DebugLog = getEnvBool("debugLog", true)
	cfg.TraceLog = getEnvBool("traceLog", true)

	mongoAuth := os.Getenv("mongoAuth")
	if mongoAuth != "" {
		cfg.MongoAuth = mongoAuth + "@"
	}

	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if str, ok := os.LookupEnv(key); ok {
		if str == "true" {
			return true
		} else if str == "false" {
			return false
		}
	}
	return fallback
}
