package config

import (
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Cookie   Cookie
	Store    Store
	Session  Session
	Metrics  Metrics
	Logger   Logger
	AWS      AWS
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string	`mapstructure:"SERVER_PORT"`
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type PostgresConfig struct {
	PostgresHost     string	`mapstructure:"POSTGRES_HOST"`
	PostgresPort     string	`mapstructure:"POSTGRES_PORT"`
	PostgresUser     string	`mapstructure:"POSTGRES_USER"`
	PostgresPassword string	`mapstructure:"POSTGRES_PASSWORD"`
	PostgresDbName   string	`mapstructure:"POSTGRES_DATABASE_NAME"`
	PostgresSSLMode  bool	`mapstructure:"POSTGRES_SSL_MODE"`
	// PgDriver         string	`mapstructure:"POSTGRES_HOST"`
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

// Cookie config
type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// Session config
type Session struct {
	Prefix string
	Name   string
	Expire int
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
}

// Store config
type Store struct {
	ImagesFolder string
}

// AWS S3
type AWS struct {
	Endpoint       string
	MinioAccessKey string
	MinioSecretKey string
	UseSSL         bool
	MinioEndpoint  string
}

// // AWS S3
// type Jaeger struct {
// 	Host        string
// 	ServiceName string
// 	LogSpans    bool
// }

// Load config file from given path
func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var cfg Config
	
	values := reflect.ValueOf(&cfg).Elem()
	 
	for i := 0; i < values.NumField(); i++ {
		err = v.Unmarshal(values.Field(i).Addr().Interface())
		if err != nil {
			log.Printf("unable to decode into struct, %v", err)
			return nil, err
		}
	}

	return &cfg, nil
}