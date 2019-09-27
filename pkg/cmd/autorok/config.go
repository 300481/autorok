package autorok

type Config struct {
	TemplateSource *TemplateSource `yaml:"templatesouce"`
	Clustername    string          `yaml:"clustername"`
}
