package config

type TencentCos struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretId   string `mapstructure:"secred-id" json:"secret-id"  yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	BaseUrl    string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"path-prefix" yaml:"path-prefix"`
}
