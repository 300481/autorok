package autorok

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Autorok is the instance type
type Autorok struct {
	Config    *config
	Templates *templates
	Cluster   *cluster
	Router    *mux.Router
}

// NewAutorok returns a new autorok instance
// configUrl must be the URL to the configuration file (YAML format)
func NewAutorok(configUrl string) *Autorok {
	// load configuration
	config := newConfig(configUrl)

	// return Autorok
	return &Autorok{
		Config:    config,
		Templates: config.TemplateSource.newTemplates(),
		Cluster:   config.newCluster(),
		Router:    mux.NewRouter(),
	}
}

// Serve runs the application in server mode
func (a *Autorok) Serve() {
	// run dummy listener to daemonize
	a.routes()
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
