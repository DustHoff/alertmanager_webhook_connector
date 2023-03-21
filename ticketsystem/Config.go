package ticketsystem

type Config struct {
	Type       string            `yaml:"type"`
	URL        string            `yaml:"url"`
	Properties map[string]string `yaml:"properties""`
}
