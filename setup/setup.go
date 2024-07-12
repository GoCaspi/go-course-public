package setup

import (
	"errors"
	"example-project/datasource"
	"example-project/model"
	"example-project/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Engine() *gin.Engine {

	databaseClient, err := datasource.NewDbClient(model.DbConfig{
		URL:      os.Getenv(DBConn),
		Database: os.Getenv(DBName),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	server.SetupService(databaseClient)
	return server.SetupEngine()
}

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	err = validateEnv()
	if err != nil {
		return err
	}

	return nil
}

const (
	DBConn        string = "DATABASE_CONNECTION_STRING"
	DBName        string = "DATABASE_NAME"
	Url           string = "URL"
	Port          string = "PORT"
	DBUrl         string = "DATABASE_URL"
	MigrationPath string = "MIGRATION_PATH"
)

func validateEnv() error {
	var err error

	vars := []string{
		DBConn,
		DBName,
		Url,
		Port,
		DBUrl,
		MigrationPath,
	}

	for i := range vars {
		err = envOk(vars[i])
		if err != nil {
			return err
		}
	}
	return err
}

func envOk(key string) error {
	if os.Getenv(key) == "" {
		return errors.New(fmt.Sprintf(".env missing value for: %s", key))
	}
	return nil
}
