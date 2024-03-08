package main

import "net/http"

func main() {
	http.HandleFunc("/", FileServer)
	http.ListenAndServe(":4000", nil)
}

func FileServer(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("", http.FileServer(http.Dir("/Users/brijesh/projects/experiments/file-server/data"))).ServeHTTP(w, r)
}
