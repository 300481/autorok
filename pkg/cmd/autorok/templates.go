package autorok

import (
	"text/template"

	"github.com/Masterminds/sprig"
)

type Templates struct {
	Ipxe    *template.Template
	Boot    *template.Template
	Install *template.Template
	RKE     *template.Template
}

type TemplateSource struct {
	Ipxe    string `yaml:"ipxe"`
	Boot    string `yaml:"boot"`
	Install string `yaml:"install"`
	RKE     string `yaml:"rke"`
}

// NewTemplates returns a new Templates instance
func newTemplates(source *TemplateSource) (*Templates, error) {
	ipxe, err := loadBytes(source.Ipxe)
	if err != nil {
		return nil, err
	}

	boot, err := loadBytes(source.Boot)
	if err != nil {
		return nil, err
	}

	install, err := loadBytes(source.Install)
	if err != nil {
		return nil, err
	}

	rke, err := loadBytes(source.RKE)
	if err != nil {
		return nil, err
	}

	return &Templates{
		Ipxe:    template.Must(template.New("ipxe").Parse(string(ipxe))).Funcs(sprig.TxtFuncMap()),
		Boot:    template.Must(template.New("boot").Parse(string(boot))).Funcs(sprig.TxtFuncMap()),
		Install: template.Must(template.New("install").Parse(string(install))).Funcs(sprig.TxtFuncMap()),
		RKE:     template.Must(template.New("rke").Parse(string(rke))).Funcs(sprig.TxtFuncMap()),
	}, nil
}
