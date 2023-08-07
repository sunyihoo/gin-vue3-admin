package config

type DsnProvider interface {
	Dsn() string
}

// Embeded 结构体可以压平到上一层，从而保持 config 文件的结构和原来一样

// GeneralDB 也被Pgsql 和 Mysql 原样使用
type GeneralDB struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址：端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // :端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Prefix       string `mapstructure:"prefix:" json:"prefix" yaml:"prefix"`                        // 全局表前缀，单独定义TableName则不生效
	Singular     string `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true便是开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine"`                         // 数据库引擎，默认为InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       string `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

type SpecializedDB struct {
	Disable   bool   `mapstructure:"disable" json:"disable" yaml:"disable"` //
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}