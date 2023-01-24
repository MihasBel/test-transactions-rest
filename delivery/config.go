package delivery

// Config configuration
type Config struct {
	Secret       string `json:"secret"`
	Address      string `json:"address"`
	StartTimeout int    `json:"start_timeout"`
	StopTimeout  int    `json:"stop_timeout"`
}
