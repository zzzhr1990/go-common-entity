package confs

// MySQLConf config for common
type MySQLConf struct {
	SlaveDatabaseConnectionString  []string `yaml:"slave-database-connection-string"`
	MasterDatabaseConnectionString []string `yaml:"master-database-connection-string"`
	DatabaseDialect                string   `yaml:"database-dialect"`
	ReadWriteSeparate              bool     `yaml:"read-write-separate"`
	AutoInitialize                 string   `yaml:"auto-initialize"`
	TableNamePrefix                string   `yaml:"table-name-prefix"`
}
