package app

import (
	"github.com/MihasBel/test-transactions-rest/adapter/broker"
	"github.com/MihasBel/test-transactions-rest/adapter/client/grpc"
	"github.com/MihasBel/test-transactions-rest/delivery"
	"github.com/MihasBel/test-transactions-rest/pkg/cache"
)

// Config exported variable to contain config values
var Config Configuration

// Configuration exported type for config
type Configuration struct {
	GRPC                  grpc.Config     `json:"grpc"`
	REST                  delivery.Config `json:"rest"`
	Redis                 cache.Config    `json:"redis"`
	Kafka                 broker.Config   `json:"kafka"`
	StartTimeout          int             `json:"start_timeout"`
	StopTimeout           int             `json:"stop_timeout"`
	ConsoleLoggingEnabled bool            `json:"console_logging_enabled"`
	FileLoggingEnabled    bool            `json:"file_logging_enabled"`
	LogDirectory          string          `json:"log_directory"`
	LogFilename           string          `json:"log_filename"`
	LogMaxBackups         int             `json:"log_max_backups"` // files
	LogMaxSize            int             `json:"log_max_size"`    // megabytes
	LogMaxAge             int             `json:"log_max_age"`     // days

}
