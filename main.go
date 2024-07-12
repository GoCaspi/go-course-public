package main

import (
	"example-project/setup"
	"log"
	"os"
)

func main() {
	err := setup.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error setting up environment: %s", err)
	}
	engine := setup.Engine()
	err = engine.Run(os.Getenv(setup.Url) + ":" + os.Getenv(setup.Port))
	if err != nil {
		log.Fatalf("Error running gin engine: %s", err)
	}
}
