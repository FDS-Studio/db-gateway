package models

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsRun    bool   `json:"isRun"`
	AutoRun  bool   `json:"autoRun"`
}
