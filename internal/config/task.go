package config

type (
	Postgres struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
	}
	HttpServer struct {
		Port int `json:"port"`
	}
	Config struct {
		Postgres   Postgres   `json:"postgres"`
		HttpServer HttpServer `json:"http_server"`
	}
)
