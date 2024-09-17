package config

type System struct {
	Addr            string          `json:"addr" yaml:"addr"`
	Database        Database        `json:"database" yaml:"database"`
	Elasticsearch   Elasticsearch   `json:"elasticsearch" yaml:"elasticsearch"`
	Alertmanger_api Alertmanger_api `json:"alertmanger_api" yaml:"alertmanger_api"`
}
