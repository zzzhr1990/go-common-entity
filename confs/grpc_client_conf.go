package confs

// GrpcClientConfig client entity
type GrpcClientConfig struct {
	DefaultEndpoint string              `yaml:"default-endpoint"`
	AutoDiscovery   AutoDiscoveryClient `yaml:"auto-discovery"`
}

// AutoDiscoveryClient AutoDiscoveryClient
type AutoDiscoveryClient struct {
	DiscoveryType   string            `yaml:"discovery-type"`
	Template        string            `yaml:"template"`
	ManualEndpoints map[string]string `yaml:"manual-endpoints"`
}
