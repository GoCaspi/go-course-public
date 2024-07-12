package setup_test

import (
	"example-project/setup"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestEngine(t *testing.T) {
	err := setup.LoadEnv("../.env")
	if err != nil {
		t.Fatal("Could not read .env")
	}
	engine := setup.Engine()
	assert.NotEqual(t, nil, engine)
}

func TestLoadEnv(t *testing.T) {
	t.Run("success loading env", func(t *testing.T) {
		defer unsetEnv()
		err := setup.LoadEnv("../.env")
		if err != nil {
			t.Fail()
		}
	})
	t.Run("error loading env", func(t *testing.T) {
		err := setup.LoadEnv(".env")
		if err == nil {
			t.Fail()
		}
	})
	t.Run("error missing value", func(t *testing.T) {
		defer unsetEnv()
		err := setup.LoadEnv("../.test.env")
		if err == nil {
			t.Fail()
		}
	})
}

func unsetEnv() {
	vars := []string{
		setup.DBConn,
		setup.DBName,
		setup.Url,
		setup.Port,
		setup.DBUrl,
		setup.MigrationPath,
	}

	for i := range vars {
		err := os.Unsetenv(vars[i])
		if err != nil {
			log.Fatal("Failure to reset ENV between tests: ", err)
		}
	}
}
