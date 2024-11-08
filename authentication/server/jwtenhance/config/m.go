package config

type M struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	MySQL struct {
		Address         string `toml:"address"`
		UserName        string `toml:"userName"`
		Password        string `toml:"password"`
		DbName          string `toml:"dbName"`
		ConnMaxLifeTime int    `toml:"connMaxLifeTime"`
		MaxIdleConn     int    `toml:"maxIdleConn"`
		MaxOpenConn     int    `toml:"maxOpenConn"`
		ConnMaxWaitTime string `toml:"connMaxWaitTime"`
	} `toml:"mysql"`
}
