package confs

// DatabaseConfig config for common
type DatabaseConfig struct {
	SlaveDatabaseConnectionString  []string `yaml:"slave-database-connection-string"`
	MasterDatabaseConnectionString []string `yaml:"master-database-connection-string"`
	DatabaseDialect                string   `yaml:"database-dialect"`
	ReadWriteSeparate              bool     `yaml:"read-write-separate"`
	AutoInitialize                 string   `yaml:"auto-initialize"`
	TableNamePrefix                string   `yaml:"table-name-prefix"`
}
