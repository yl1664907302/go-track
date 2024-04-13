package config

type Database struct {
	MysqlHost         string `json:"mysqlhost" yaml:"host"`
	MysqlPort         int    `json:"mysqlport" yaml:"port"`
	MysqlUser         string `json:"mysqluser" yaml:"user"`
	MysqlPassword     string `json:"mysqlpassword" yaml:"password"`
	MysqlDatabasename string `json:"mysqldatabasename" yaml:"databasename"`
}
