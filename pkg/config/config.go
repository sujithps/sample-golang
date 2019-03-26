package config

import (
	"github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

var appConfig *Config

type Config struct {
	port            int
	log             LogConfig
	newrelic        newrelic.Config
	newrelicEnabled bool
	mongoURL        string
	mongoDbName     string
}

func Load() {
	viper.SetDefault("APP_PORT", "3000")
	viper.SetDefault("LOG_LEVEL", "warn")
	viper.SetDefault("LOG_FORMAT", "json")
	viper.SetDefault("APP_NAME", "sample-golang")
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	appConfig = &Config{
		port: mustGetInt("APP_PORT"),
		log: LogConfig{
			logLevel: mustGetString("LOG_LEVEL"),
			format:   mustGetString("LOG_FORMAT"),
		},
		newrelic:        newrelicConfig(),
		newrelicEnabled: mustGetBool("NEW_RELIC_ENABLED"),
		mongoURL:        mustGetString("MONGO_DB_URL"),
		mongoDbName:     mustGetString("MONGO_DB_NAME"),
	}
}

func Log() LogConfig {
	return appConfig.log
}

func Port() int {
	return appConfig.port
}

func NewRelic() newrelic.Config {
	return appConfig.newrelic
}

func MongoURL() string {
	return appConfig.mongoURL
}

func MongoDBName() string {
	return appConfig.mongoDbName
}
