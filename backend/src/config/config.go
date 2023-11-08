package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   LoggerConfig
	Otp      OtpConfig
	JWT      JWTConfig
	Cors     CorseConfig
	Password PasswordConf
}
type CorseConfig struct {
	AllowOrigins string
}
type ServerConfig struct {
	InternalPort int
	ExternalPort int
	RunMode      string
}

type JWTConfig struct {
	Secret                     string
	RefreshSecret              string
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

type PostgresConfig struct {
	Host            string
	User            string
	Password        string
	DbName          string
	SslMode         string
	Port            int
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               int
	Password           string
	Db                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        int
	IdleCheckFrequency int
}
type OtpConfig struct {
	Digits     int
	ExpireTime time.Duration
	Limiter    time.Duration
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}
type PasswordConf struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

func configPath(env string) string {
	if env == "production" {
		return "../config/config-production.yml"
	}
	if env == "docker" {
		return "../config/config-docker.yml"
	}
	return "../config/config-development.yml"
}

func loadConfig(fileName, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("cfg path is incorrect")
		}
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (cfg *Config, err error) {
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfig() *Config {
	env := os.Getenv("APP_ENV")
	// TODO: fix environ
	path := configPath(env)
	v, err := loadConfig(path, "yml")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
