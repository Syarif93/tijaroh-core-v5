package process

import (
	"log"
	"os"
	"time"
)

func LoggerFile() *os.File {
	loc, timeErr := time.LoadLocation(Env().TimeZone)
	if timeErr != nil {
		panic(timeErr)
	}
	time := time.Now().In(loc)

	var fileName string = time.Format("2006-01-02 15:04:05")

	if Env().AppEnv == "development" {
		fileName = "dev"
	}

	file, err := os.OpenFile("./logs/"+fileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
