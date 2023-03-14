package config

type GRPC struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func (g *GRPC) GetAddr() string {
	return g.Host + ":" + g.Port
}
