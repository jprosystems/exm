package model

type DbServer struct {
	Host    string `json:"DB_HOST"`
	User    string `json:"DB_USER"`
	Pass    string `json:"DB_PASS"`
	DbPath  string `json:"DB_PATH"`
	DbRdbms string `json:"DB_RDBMS"`
}
