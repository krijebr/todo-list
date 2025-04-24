package config

type ConfigPostgres struct {
	Host     string `json:"host"`
	PortDB   int    `json:"port_db"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}
type ConfigRouter struct {
	PortR int `json:"port_r"`
}
type Config struct {
	ConfigPostgres
	ConfigRouter
}
