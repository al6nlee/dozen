package config

type PostgresConfig struct {
	Host     string `destructure:"host" json:"host" yaml:"host"`
	Port     string `destructure:"port" json:"port" yaml:"port"`
	User     string `destructure:"user" json:"user" yaml:"user"`
	Password string `destructure:"password" json:"password" yaml:"password"`
	DBName   string `destructure:"dbname" json:"dbname" yaml:"dbname"`
}

type RedisConfig struct {
	Host     string `destructure:"host" json:"host" yaml:"host"`
	Port     int    `destructure:"port" json:"port" yaml:"port"`
	Password string `destructure:"password" json:"password" yaml:"password"`
}

type ServerConfig struct {
	Name     string         `destructure:"name" json:"name" yaml:"name"`
	Host     string         `destructure:"host" json:"host" yaml:"host"`
	PORT     int            `destructure:"port" json:"port" yaml:"port"`
	Mode     string         `destructure:"mode" json:"mode" yaml:"mode"`
	Postgres PostgresConfig `destructure:"postgres" json:"postgres" yaml:"postgres"`
	Redis    RedisConfig    `destructure:"redis" json:"redis" yaml:"redis"`
}
