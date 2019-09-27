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
	// load configuration
	config := Config{}
	err := loadObject(
		configUrl,
		YAML,
		&config,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// load templates
	templates, err := newTemplates(config.TemplateSource)
	if err != nil {
		log.Fatalln(err)
	}

	return &Autorok{
		Config:    &config,
		Templates: templates,
	}
}

// Serve runs the application in server mode
func (a *Autorok) Serve() {
	// run dummy listener to daemonize
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
