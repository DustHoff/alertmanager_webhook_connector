package ticketsystem

type Config struct {
	Type     string `yaml:"type"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Target   string `yaml:"target"`
}
