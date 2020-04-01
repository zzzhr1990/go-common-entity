package confs

// RDBMSConf config for common not just for mysql
type RDBMSConf struct {
	SlaveConnectionString  []string `yaml:"slave-connection-string"`
	MasterConnectionString []string `yaml:"master-connection-string"`
	Dialect                string   `yaml:"dialect"`
	ReadWriteSeparate      bool     `yaml:"read-write-separate"`
	AutoInitialize         bool     `yaml:"auto-initialize"`
	TableNamePrefix        string   `yaml:"table-name-prefix"`
}

// RDBMSMultiMasterDbConf config for common not just for mysql
type RDBMSMultiMasterDbConf struct {
	// NamePrefix []string `yaml:"name-prefix"`
	// MasterConnectionString []string `yaml:"master-connection-string"`
	Dbs             map[string]string `yaml:"dbs"`
	Dialect         string            `yaml:"dialect"`
	AutoInitialize  bool              `yaml:"auto-initialize"`
	TableNamePrefix string            `yaml:"table-name-prefix"`
}

// RDBMSMultiSlaveDbConf config for common not just for mysql
type RDBMSMultiSlaveDbConf struct {
	// NamePrefix []string `yaml:"name-prefix"`
	// MasterConnectionString []string `yaml:"master-connection-string"`
	Dbs             map[string][]string `yaml:"dbs"`
	Dialect         string              `yaml:"dialect"`
	TableNamePrefix string              `yaml:"table-name-prefix"`
}
