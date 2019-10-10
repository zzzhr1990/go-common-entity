package confs

// MySQLConf config for common
type MySQLConf struct {
	SlaveConnectionString  []string `yaml:"slave-connection-string"`
	MasterConnectionString []string `yaml:"master-connection-string"`
	Dialect                string   `yaml:"dialect"`
	ReadWriteSeparate      bool     `yaml:"read-write-separate"`
	AutoInitialize         string   `yaml:"auto-initialize"`
	TableNamePrefix        string   `yaml:"table-name-prefix"`
}
