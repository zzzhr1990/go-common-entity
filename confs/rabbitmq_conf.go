package confs

// QcloudCMQConfig config CMQ setting
type QcloudCMQConfig struct {
	SecretID  string `yaml:"secret-id"`
	SecretKey string `yaml:"secret-key"`
	QueueName string `yaml:"queue"`
}

// AmqpConfig config MQ setting
type AmqpConfig struct {
	ConnectString    string         `yaml:"connect-string"`
	Name             string         `yaml:"name"`
	Exchange         ExchangeConfig `yaml:"exchange"`
	Bind             BindConfig     `yaml:"bind"`
	Durable          bool           `yaml:"durable"`
	DeleteWhenUnused bool           `yaml:"delete-when-unused"`
	Exclusive        bool           `yaml:"exclusive"`
	NoWait           bool           `yaml:"no-wait"`
	RoutingKey       string         `yaml:"routing-key"`
}

// ExchangeConfig exchange
type ExchangeConfig struct {
	Name             string `yaml:"name"`
	Durable          bool   `yaml:"durable"`
	DeleteWhenUnused bool   `yaml:"delete-when-unused"`
	Exclusive        bool   `yaml:"exclusive"`
	NoWait           bool   `yaml:"no-wait"`
	Internal         bool   `yaml:"internal"`
	Kind             string `yaml:"kind"`
}

// ExchangeBind(destination string, key string, source string, noWait bool

// BindConfig cfg
type BindConfig struct {
	Destination string `yaml:"destination"`
	Key         string `yaml:"key"`
	NoWait      bool   `yaml:"no-wait"`
	Source      string `yaml:"source"`
}
