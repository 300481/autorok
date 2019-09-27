package autorok

import (
	"log"
	"text/template"

	"github.com/Masterminds/sprig"
)

type templates struct {
	Ipxe    *template.Template
	Boot    *template.Template
	Install *template.Template
	RKE     *template.Template
}

type templateSource struct {
	Ipxe    string `yaml:"ipxe"`
	Boot    string `yaml:"boot"`
	Install string `yaml:"install"`
	RKE     string `yaml:"rke"`
}

// NewTemplates returns a templates instance
func (s *templateSource) newTemplates() *templates {
	ipxe, err := loadBytes(s.Ipxe)
	if err != nil {
		log.Fatalln(err)
	}

	boot, err := loadBytes(s.Boot)
	if err != nil {
		log.Fatalln(err)
	}

	install, err := loadBytes(s.Install)
	if err != nil {
		log.Fatalln(err)
	}

	rke, err := loadBytes(s.RKE)
	if err != nil {
		log.Fatalln(err)
	}

	return &templates{
		Ipxe:    template.Must(template.New("ipxe").Parse(string(ipxe))).Funcs(sprig.TxtFuncMap()),
		Boot:    template.Must(template.New("boot").Parse(string(boot))).Funcs(sprig.TxtFuncMap()),
		Install: template.Must(template.New("install").Parse(string(install))).Funcs(sprig.TxtFuncMap()),
		RKE:     template.Must(template.New("rke").Parse(string(rke))).Funcs(sprig.TxtFuncMap()),
	}
}
