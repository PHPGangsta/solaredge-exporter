package config

import (
	"os"
	"strconv"

	cfg "github.com/infinityworks/go-common/config"
	)

// Config struct holds all of the runtime configuration for the application
type Config struct {
	*cfg.BaseConfig
	Site          int64
	APIToken      string
	Timezone      string
}

func Init() Config {
	ac := cfg.Init()
	siteID, err := strconv.ParseInt(os.Getenv("SITE_ID"), 10, 64)
	if err != nil {
		panic("Can't get site ID")
	}
	APIToken := os.Getenv("API_TOKEN")

	// Set default timezone to CST (America/Chicago) for backward compatibility
	timezone := os.Getenv("TIMEZONE")
	if timezone == "" {
		timezone = "America/Chicago"
	}

	appConfig := Config{
		&ac,
		siteID,
		APIToken,
		timezone,
	}

	return appConfig
}
