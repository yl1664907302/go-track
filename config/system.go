package config

type System struct {
	Addr     string   `json:"addr" yaml:"addr"`
	Database Database `json:"database" yaml:"database"`
}
