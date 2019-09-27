package autorok

type Config struct {
	TemplateSource TemplateSource `yaml:"templatesource"`
	Clustername    string         `yaml:"clustername"`
}
