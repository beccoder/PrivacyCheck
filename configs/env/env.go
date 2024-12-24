package env

import "sync"

var envProject EnvProject

type EnvProject struct {
	Mode        string `env:"MODE"`
	HttpPort    int    `env:"HTTP_PORT"`
	StoragePath string `env:"STORAGE_PATH"`

	PgDB       string `env:"POSTGRES_DB"`
	PgHost     string `env:"POSTGRES_HOST"`
	PgPort     uint   `env:"POSTGRES_PORT"`
	PgUser     string `env:"POSTGRES_USER"`
	PgPassword string `env:"POSTGRES_PASSWORD"`

	JwtSecret  string `env:"JWT_SECRET"`
	JwtExpire  int64  `env:"JWT_EXPIRE"`
	JwtRefresh int64  `env:"JWT_REFRESH"`
}

func init() {
	sync.OnceFunc(func() {
		MustLoad(&envProject)
	})()
}

func ProjectEnv() EnvProject {
	return envProject
}
