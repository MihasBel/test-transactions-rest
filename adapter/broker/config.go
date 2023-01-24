package broker

// Config configuration
type Config struct {
	KafkaURL        string `json:"kafka_url"`
	PartitionsCount int    `json:"partitions_count"`
}
