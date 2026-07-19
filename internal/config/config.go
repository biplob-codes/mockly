package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)
type Config struct{
	Env string
	DatabaseName string
	Port string
}
func Load() *Config{
	if err:=godotenv.Load();err!=nil{
		fmt.Println("no .env file found, using real environmental variables")
	}
	env:=os.Getenv("ENV")
	if env==""{
		env="local"
	}
	databaseName:=os.Getenv("DATABASE_NAME")
	if databaseName==""{
		databaseName="mockly.db"
	}
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}

	return &Config{Env: env,DatabaseName: databaseName,Port: port}
	
}