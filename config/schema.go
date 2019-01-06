package config

// ServiceConfiguration is the top level configuration struct which is loaded from the defined source(s)
type ServiceConfiguration struct {
	Log LogConfiguration `json:"log"`
	Environment string `json:"environment"`
	Database DatabaseConfiguration `json:"database"`
	Metrics MetricsConfiguration `json:"metrics"`
}

type LogConfiguration struct {
	Level           string `json:"level"`
	Format          string `json:"format"` // json or text
}

type DatabaseConfiguration struct {
	Uri string `json:"uri"`
	Dialect string `json:"dialect"`
}

type MetricsConfiguration struct {
	Endpoint string `json:"endpoint"`
	Port int `json:"port"`
}
