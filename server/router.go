package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	// Single Page Application files handler
	spa := spaHandler{staticPath: "server/client/build/", indexPath: "index.html"}
	spaHandler := http.StripPrefix("/client", spa)
	r.
		Methods("GET").
		PathPrefix("/client").
		Name("SPA").
		Handler(spaHandler)

	return r
}

// spaHandler implements the http.Handler interface, so we can use it to respond to HTTP requests.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir on the SPA handler. If a file is found,
// it will be served. If not, the file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal, if we fail send 400 and stop
	path, err := filepath.Abs("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(path, h.staticPath)
	path = filepath.Join(path, r.URL.Path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error stating the file, send 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
