package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// set hostname (used for demo)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error:", err)
	}

	data := struct {
		Title    string
		Hostname string
	}{
		Title:    "Kubernetes Pod Load Balancer Demo (refresh page)",
		Hostname: hostname,
	}

	t, err := template.New("index.html").ParseFiles("index.html")

	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	err = t.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
	}

}

// used to dump headers for debugging
func debugHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	// disable cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// set hostname (used for demo)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error:", err)
	}

	// dump headers
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%v", string(requestDump))
	fmt.Fprintf(w, "Served-By: %v\n", hostname)
	fmt.Fprintf(w, "Serving-Time: %s", time.Now().Sub(startTime))
	return

}

// mux
var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/debug", debugHandler)
	http.Handle("/", router)

	fmt.Println("Listening on port 5005...")
	http.ListenAndServe(":5005", handlers.CompressHandler(router))

}
