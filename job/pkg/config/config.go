package config

var cfg *Cfg

type Cfg struct {
	Postgres struct {
		Dsn string
	}
}

func Set(c *Cfg) {
	cfg = c
}

func Get() *Cfg {
	return cfg
}
