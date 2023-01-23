package grpc

type Config struct {
	Endpoint         string `json:"endpoint" `
	DealTimeout      int    `json:"deal_timeout" `
	KeepAliveTime    int    `json:"keep_alive_time" `
	KeepAliveTimeout int    `json:"keep_alive_timeout" `
}
