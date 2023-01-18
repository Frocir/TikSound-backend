package config

type Log struct {
	OutPut      []string `json:"outPut" mapstructure:"output"`
	ErrorOutPut []string `json:"errorOutPut" mapstructure:"errorOutput"`
	Level       string   `json:"level" mapstructure:"level"`
	Format      string   `json:"format" mapstructure:"format"`
	Name        string   `json:"name" mapstructure:"name"`
}
