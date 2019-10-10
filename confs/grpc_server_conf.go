package confs

// GrpcServerConfig server config
type GrpcServerConfig struct {
	Port          int    `yaml:"port"`
	Endpoint      string `yaml:"endpoint"`
	AutoDiscovery struct {
		DiscoveryType string
		Template      string
		ServerAddress string
	} `yaml:"auto-discovery"`
}
