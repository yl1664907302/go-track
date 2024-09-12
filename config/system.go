package config

type System struct {
	Addr             string           `json:"addr" yaml:"addr"`
	Database         Database         `json:"database" yaml:"database"`
	Elasticsearch    Elasticsearch    `json:"elasticsearch" yaml:"elasticsearch"`
	Alertmanager_api Alertmanager_api `json:"alertmanager_api" yaml:"alertmanager_api"`
}
