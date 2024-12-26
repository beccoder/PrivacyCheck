package env

import (
	"sync"
)

var envProject EnvProject

type EnvProject struct {
	HttpPort int `env:"HTTP_PORT"`

	PgDB       string `env:"POSTGRES_DB"`
	PgHost     string `env:"POSTGRES_HOST"`
	PgPort     uint   `env:"POSTGRES_PORT"`
	PgUser     string `env:"POSTGRES_USER"`
	PgPassword string `env:"POSTGRES_PASSWORD"`

	JwtSecret string `env:"JWT_SECRET"`
	JwtExpire int    `env:"JWT_EXPIRE"`
}

func init() {
	sync.OnceFunc(func() {
		MustLoad(&envProject)
	})()
}

func ProjectEnv() EnvProject {
	return envProject
}
