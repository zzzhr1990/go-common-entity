package confs

// KafkaConsumerConfig support qc
type KafkaConsumerConfig struct {
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	Mechanism string   `yaml:"mechanism"`
	GroupID   string   `yaml:"group-id"`
	Topic     string   `yaml:"topic"`
	Partition int      `yaml:"partition"`
	Brokers   []string `yaml:"brokers"`
}

// KafkaProducerConfig support qc
type KafkaProducerConfig struct {
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	Mechanism string   `yaml:"mechanism"`
	Topic     string   `yaml:"topic"`
	Async     bool     `yaml:"async"`
	Brokers   []string `yaml:"brokers"`
}
