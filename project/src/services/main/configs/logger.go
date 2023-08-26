package configs

import (
	"os"
	"project/src/pkg/utils"
	"time"
)

func GinLogger() (*os.File, error) {
	utils := utils.New()
	path, err := utils.GetFilePath(&[]string{"src", "services", "main", "configs", "logs"})
	if err != nil {
		return nil, err
	}
	today := time.Now().Format("2006-01-02")

	osDirectorySpliter := os.PathSeparator

	f, err := os.Create(*path + string(osDirectorySpliter) + today + ".log")
	if err != nil {
		return nil, err
	}

	return f, nil
}
