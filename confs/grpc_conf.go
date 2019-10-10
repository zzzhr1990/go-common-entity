package confs

// GrpcConfig common config
type GrpcConfig struct {
	Server GrpcServerConfig `yaml:"server"`
	Client GrpcClientConfig `yaml:"client"`
}
