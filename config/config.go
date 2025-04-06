package config

import (
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/joho/godotenv"
)

type Env struct {
	MONGODB_URL    string
	DATABASE       string
	ENVIROMENT     string
	PORT           string
	JWT_SECRET_KEY string
	BASE_URL       string
}

func GetConfigPath() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	
	return filepath.Join(root, ".env"), nil
}

func (ev *Env) setupVariables() {
	file, err :=GetConfigPath()
	if err != nil {
		log.Fatalf("Error finding .env file %v\n", err)
	}

	err = godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading .env %v file %v\n",file, err)
	}

	variables := []string{"MONGODB_URL", "DATABASE", "ENVIROMENT", "PORT", "JWT_SECRET_KEY", "BASE_URL"}
	val := reflect.ValueOf(ev).Elem()

	for _, v := range variables {
		if field := val.FieldByName(v); field.IsValid() {
			if data, exists := os.LookupEnv(v); exists {
				field.SetString(data)
			}
		}
	}
}

func getEnv() Env {

	env := Env{}
	env.setupVariables()

	return env

}

var EnvData = getEnv()
