package confs

// RedisConfig config Redis setting
type RedisConfig struct {
	Network  string `yaml:"network"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	ReadOnly bool   `yaml:"read-only"`
}
