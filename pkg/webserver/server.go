package webserver

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/madacluster/gosero/pkg/relay"
)

var statusCaldera = "stop"

//go:embed public
var embeddedFiles embed.FS

func start(w http.ResponseWriter, r *http.Request) {
	print("Start!")
	statusCaldera = "start"
	relay.Start()
	w.Write([]byte(statusCaldera))
}

func stop(w http.ResponseWriter, r *http.Request) {
	print("Stop!")
	statusCaldera = "stop"
	relay.Stop()
	w.Write([]byte(statusCaldera))
}

func status(w http.ResponseWriter, r *http.Request) {
	print("Status!")
	w.Write([]byte(statusCaldera))
}

func Webserver() {
	http.HandleFunc("/start", start)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/status", status)
	// u, _ := url.Parse("http://localhost:8080")
	// http.Handle("/", httputil.NewSingleHostReverseProxy(u))
	fsys, err := fs.Sub(embeddedFiles, "public")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":8081", nil)
}
