package confs

// MongoConf config
type MongoConf struct {
	Dialect        string `yaml:"dialect"`
	ConnectString  string `yaml:"connect-string"`
	DatabaseName   string `yaml:"database-name"`
	CollectionName string `yaml:"collection-name"`
}
