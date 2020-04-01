package confs

// MySQLConf config for common just for mysql
type MySQLConf struct {
	SlaveConnectionString  []string `yaml:"slave-connection-string"`
	MasterConnectionString []string `yaml:"master-connection-string"`
	Dialect                string   `yaml:"dialect"`
	ReadWriteSeparate      bool     `yaml:"read-write-separate"`
	AutoInitialize         bool     `yaml:"auto-initialize"`
	TableNamePrefix        string   `yaml:"table-name-prefix"`
}

// MySQLMultiDbConf config for common just for mysql
type MySQLMultiDbConf struct {
	NamePrefix []string `yaml:"name-prefix"`
	// MasterConnectionString []string `yaml:"master-connection-string"`
	Dbs             map[string]string `yaml:"dbs"`
	Dialect         string            `yaml:"dialect"`
	AutoInitialize  bool              `yaml:"auto-initialize"`
	TableNamePrefix string            `yaml:"table-name-prefix"`
}
