package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init(file ...string) {
	if len(file) == 0 {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return
	}
	err := godotenv.Load(file[0])
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetString(env string) string {
	return os.Getenv(env)
}

func GetInt(env string) int {
	envInt, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		log.Fatal("Error loading env variable")
	}
	return envInt
}

func GetFloat(env string) float64 {
	envFloat, err := strconv.ParseFloat(env, 64)
	if err != nil {
		log.Fatal("Error loading env variable")
	}
	return envFloat
}
