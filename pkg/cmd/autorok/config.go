package autorok

import "log"

type config struct {
	TemplateSource templateSource `yaml:"templatesource"`
	ClusterConfig  string         `yaml:"clusterconfig"`
	BootServer     string         `yaml:"bootserver"`
	ClusterName    string         `yaml:"clustername"`
	NodeCount      int            `yaml:"nodecount"`
	PublicKey      string         `yaml:"publickkey"`
	StartCIDR      string         `yaml:"startcidr"`
	Gateway        string         `yaml:"gateway"`
	MTU            int            `yaml:"mtu"`
	DHCP           bool           `yaml:"dhcp"`
	Nameservers    []string       `yaml:"nameservers"`
}

func newConfig(configUrl string) *config {
	c := &config{}
	err := loadObject(configUrl, YAML, c)
	if err != nil {
		log.Fatalln(err)
	}
	return c
}
