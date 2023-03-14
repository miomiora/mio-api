package config

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	Timeout  string `yaml:"timeout"`
	MaxConn  int    `yaml:"maxConn"`
	MaxOpen  int    `yaml:"maxOpen"`
}

func (m *Mysql) GetDsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local&timeout=" + m.Timeout
}
