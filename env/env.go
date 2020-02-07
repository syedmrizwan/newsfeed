package env

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
)

type envFile struct {
	DbName     string
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	BuildEnv   string
	ServerPort string
	DbPoolSize int
}

func (e *envFile) GetAddr() string {
	return e.DbHost + ":" + e.DbPort
}

var Env *envFile

func init() {
	_ = godotenv.Load()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	DbPoolSize, _ := strconv.Atoi(os.Getenv("DB_POOL_SIZE"))

	Env = &envFile{
		DbName:     os.Getenv("DB_NAME"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		BuildEnv:   os.Getenv("BUILD_ENV"),
		ServerPort: os.Getenv("SERVER_PORT"),
		DbPoolSize: DbPoolSize,
	}
}
