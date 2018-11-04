package config

// ServiceConfiguration is the top level configuration struct which is loaded from the defined source(s)
type ServiceConfiguration struct {
	Log LogConfiguration
}

type LogConfiguration struct {
	Level           string `json:"level"`
	Format          string `json:"format"` // json or text
}
