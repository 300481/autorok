package autorok

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Autorok is the instance type
type Autorok struct {
	Config    *Config
	Templates *Templates
}

// NewAutorok returns a new autorok instance
// configUrl must be the URL to the configuration file (YAML format)
func NewAutorok(configUrl string) *Autorok {
	config := &Config{}
	err := loadObject(
		configUrl,
		YAML,
		config,
	)
	if err != nil {
		log.Fatalln(err)
	}

	templates, err := NewTemplates(config.TemplateSource)
	if err != nil {
		log.Fatalln(err)
	}

	return &Autorok{
		Config:    config,
		Templates: templates,
	}
}

// Execute runs the application
func (a *Autorok) Execute() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
