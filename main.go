package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var files map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(files[r.RequestURI])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, string(b))
}

func main() {
	cfg, err := readConfig(filepath.Join(os.Getenv("HOME"), ".logcat"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range cfg.Logs {
		key := fmt.Sprintf("/%s", k)
		files[key] = v.Path
		http.HandleFunc(key, handler)
	}

	fmt.Println(files)

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func init() {
	files = make(map[string]string)
}
