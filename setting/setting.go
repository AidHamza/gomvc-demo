package setting

import (
	"os"

	"github.com/AidHamza/gomvc-demo/models/logger"
)

// option .
type option struct {
	MongoURL    string
	MongoDBName string
}

// Option .
var Option *option

func init() {
	logger.Log("setting#init")

	if os.Getenv("GO_ENV") == "development" {
		Option = optionDev

		logger.Log("setting#init: development")
		return
	}

	logger.Log("setting#init: production")
	Option = optionPro
}
