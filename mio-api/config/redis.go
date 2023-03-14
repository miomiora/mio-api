package config

type Redis struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

func (r *Redis) GetAddr() string {
	return r.Host + ":" + r.Port
}
