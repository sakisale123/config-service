package config

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Configuration struct {
	ID         string         `json:"id"`
	Version    string         `json:"version"`
	Parameters []KeyValuePair `json:"parameters"`
}

type ConfigurationGroup struct {
	ID             string          `json:"id"`
	Version        string          `json:"version"`
	Configurations []Configuration `json:"configurations"`
}
