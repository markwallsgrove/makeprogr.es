package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name     string
	Email    string
	Username string
}

type Home struct {
	staticFiles http.FileSystem
}

func (home *Home) homeHandler(resp http.ResponseWriter, req *http.Request) {
	httpFile, err := home.staticFiles.Open("/static/main.html")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Could not find http file"))
		return
	}

	fmt.Println("hey toastcat!")
	content, err := ioutil.ReadAll(httpFile)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Could not read main.html"))
		log.Printf(err.Error())
		return
	}

	resp.Write(content)
}

// TODO: not sure this will cause a performance issue
// Home is copied each time this call is made
func (home Home) Open(name string) (http.File, error) {
	fmt.Println("opening", name)
	return home.staticFiles.Open(name)
}

func main() {
	home := Home{staticFiles: FS(false)}
	r := mux.NewRouter()
	r.HandleFunc("/", home.homeHandler)
	r.PathPrefix("/static/").Handler(http.FileServer(home))
	http.ListenAndServe("0.0.0.0:8080", r)
}
