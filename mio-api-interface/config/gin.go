package config

type Gin struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func (g *Gin) GetAddr() string {
	return g.Host + ":" + g.Port
}
