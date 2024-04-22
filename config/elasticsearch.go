package config

type Elasticsearch struct {
	Eshost   string `json:"eshost"yaml:"eshost"`
	Esport   string `json:"esport"yaml:"esport"`
	Username string `json:"username"yaml:"username"`
	Password string `json:"password"yaml:"password"`
}
