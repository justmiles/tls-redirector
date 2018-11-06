package main

import "net/http"
import "log"
import "os"

func main() {
	var listenPort = port()
	log.Printf("Redirector starting on port %s", listenPort)
	if err := http.ListenAndServe(":"+listenPort, http.HandlerFunc(redirect)); err != nil {
		log.Fatalf("Could not listen on %s\n: %v", listenPort, err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

// set default port
func port() string {
	if len(os.Getenv("LISTEN_PORT")) != 0 {
		return os.Getenv("LISTEN_PORT")
	}
	return "80"
}
