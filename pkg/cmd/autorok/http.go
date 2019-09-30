package autorok

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *Autorok) routes() {
	a.Router.HandleFunc("/ipxe", a.ipxeHandler)
	a.Router.HandleFunc("/boot", a.bootHandler)
	a.Router.HandleFunc("/install/{uuid}", a.installHandler)
	a.Router.HandleFunc("/rke", a.rkeHandler)
}

func (a *Autorok) ipxeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ipxe.handler:", r.Method, "request from ", r.RemoteAddr)
	a.Templates.Ipxe.Execute(w, a.Config)
}

func (a *Autorok) bootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("boot.handler:", r.Method, "request from ", r.RemoteAddr)
	a.Templates.Boot.Execute(w, a.Config)
}

func (a *Autorok) installHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("install.handler:", r.Method, "request from ", r.RemoteAddr)

	uuid := mux.Vars(r)["uuid"]
	n := a.getNode(uuid)
	if n == nil {
		log.Println("install.handler: max cluster size reached")
		io.WriteString(w, "")
		return
	}
	a.Templates.Install.Execute(w, n)
}

func (a *Autorok) rkeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("rke.handler:", r.Method, "request from ", r.RemoteAddr)
	a.Templates.RKE.Execute(w, a.Cluster)
}
