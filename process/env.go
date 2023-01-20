package process

import (
	"os"
)

type envModel struct {
	AppEnv, AppPort, TimeZone string
}

func Env() envModel {
	env := envModel{}
	env.AppEnv = os.Getenv("APP_ENV")
	env.AppPort = os.Getenv("APP_PORT")
	env.TimeZone = os.Getenv("TIME_ZONE")

	return env
}
