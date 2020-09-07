/*
 * Logcat - a small utility useful for exposing logs to http requests.
 * Copyright (C) 2020  Nicolò Santamaria
 *
 * Logcat is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Logcat is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with logcat.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var files = make(map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(files[r.RequestURI])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, string(b))
}

func main() {
	var home = os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("UserProfile")
	}

	cfg, err := readConfig(filepath.Join(home, ".logcat"))
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range cfg.Logs {
		key := fmt.Sprintf("/%s", k)
		files[key] = v.Path
		http.HandleFunc(key, handler)
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}
