package config

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket " `
	ImgPath       string `mapstructure:"img-path" json:"img-path" yaml:"imgPath"`
	UseHttps      bool   `mapstructure:"user-https" json:"use-https" yaml:"useHttps"`
	AccessKey     string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn=dimains" json:"use-cdn-domains" yaml:"use-cdn-domains"`
}
