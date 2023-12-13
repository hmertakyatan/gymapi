package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
}

type TokenConfig struct {
	AccessTokenExpiredIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	AccessTokenSecret     string        `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenExpiredIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	RefreshTokenMaxAge    int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
	RefreshTokenSecret    string        `mapstructure:"REFRESH_TOKEN"`
}

type ServerConfig struct {
	CustomerServerPort         string `mapstructute:"CUSTOMER_SERVER_PORT"`
	MembershipServerPort       string `mapstructute:"MEMBERSHIP_SERVER_PORT"`
	MembershipTypeServerPort   string `mapstructute:"MEMBERSHIP_TYPE_SERVER_PORT"`
	PaymentServerPort          string `mapstructute:"PAYMENT_SERVER_PORT"`
	PersonalTrainingServerPort string `mapstructute:"PERSONAL_TRAINING_SERVER_PORT"`
	PersonnelServerPort        string `mapstructute:"PERSONNEL_SERVER_PORT"`
	AuthServerPort             string `mapstructute:"AUTH_SERVER_PORT"`
	UserServerPort             string `mapstructute:"USER_SERVER_PORT"`
}

type Config struct {
	DBConfig
	ServerConfig
	TokenConfig
}

//Config initializer

func LoadConfig(path string) (Config, error) {
	var config Config

	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file:", err)
		return config, err
	}

	config.DBHost = os.Getenv("DB_HOST")
	config.DBUsername = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASS")
	config.DBName = os.Getenv("DB_NAME")
	config.DBPort = os.Getenv("DB_PORT")
	config.CustomerServerPort = os.Getenv("CUSTOMER_SERVER_PORT")
	config.MembershipServerPort = os.Getenv("MEMBERSHIP_SERVER_PORT")
	config.MembershipTypeServerPort = os.Getenv("MEMBERSHIP_TYPE_SERVER_PORT")
	config.PaymentServerPort = os.Getenv("PAYMENT_SERVER_PORT")
	config.PersonalTrainingServerPort = os.Getenv("PERSONAL_TRAINING_SERVER_PORT")
	config.PersonnelServerPort = os.Getenv("PERSONNEL_SERVER_PORT")
	config.AuthServerPort = os.Getenv("AUTH_SERVER_PORT")
	config.UserServerPort = os.Getenv("USER_SERVER_PORT")

	accessTokenExpiredInStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	accessTokenExpireDuration, err := time.ParseDuration(accessTokenExpiredInStr)
	if err != nil {
		log.Println("Error parsing ACCESS_TOKEN_EXPIRED_IN:", err)
		return config, err
	}

	accessTokenMaxAgeStr := os.Getenv("ACCESS_TOKEN_MAXAGE")
	accessTokenMaxAge, err := strconv.Atoi(accessTokenMaxAgeStr)
	if err != nil {
		log.Println("Error parsing ACCESS_TOKEN_MAXAGE")
		return config, err
	}

	config.AccessTokenExpiredIn = accessTokenExpireDuration
	config.AccessTokenMaxAge = accessTokenMaxAge
	config.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

	refreshTokenExpiredInStr := os.Getenv("REFRESH_TOKEN_EXPIRED_IN")
	refreshTokenExpireDuration, err := time.ParseDuration(refreshTokenExpiredInStr)
	if err != nil {
		log.Println("Error parsing TOKEN_EXPIRED_IN:", err)
		return config, err
	}

	refreshTokenMaxAgeStr := os.Getenv("REFRESH_TOKEN_MAXAGE")
	refreshTokenMaxAge, err := strconv.Atoi(refreshTokenMaxAgeStr)
	if err != nil {
		log.Println("Error parsing REFRESH_TOKEN_MAXAGE")
		return config, err
	}

	config.RefreshTokenExpiredIn = refreshTokenExpireDuration
	config.RefreshTokenMaxAge = refreshTokenMaxAge
	config.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")

	return config, nil
}
