package config

// ServiceConfiguration is the top level configuration struct which is loaded from the defined source(s)
type ServiceConfiguration struct {
	Log LogConfiguration
	LocalEnv bool `json:"local_env"`
	Database DatabaseConfiguration
}

type LogConfiguration struct {
	Level           string `json:"level"`
	Format          string `json:"format"` // json or text
}

type DatabaseConfiguration struct {
	Uri string `json:"uri"`
	Dialect string `json:"dialect"`
}
