package app

import "github.com/MihasBel/test-transactions-rest/adapter/client/grpc"

// Config exported variable to contain config values
var Config Configuration

// Configuration exported type for config
type Configuration struct {
	GRPC                  grpc.Config `json:"grpc"`
	ConnectionString      string
	Database              string
	Collection            string
	APIKey                string
	Address               string
	StartTimeout          int
	StopTimeout           int
	ConsoleLoggingEnabled bool
	FileLoggingEnabled    bool
	LogDirectory          string
	LogFilename           string
	LogMaxBackups         int // files
	LogMaxSize            int // megabytes
	LogMaxAge             int // days
	KafkaURL              string
	PartitionsCount       int
	RedisURL              string
	RedisPass             string
}
