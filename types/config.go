package types

// Config ...
type Config struct {
	Port int `mapstructure:"port"`
	DB   db  `mapstructure:"db"`
}

type db struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
