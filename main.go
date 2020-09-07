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
	"log"
	"net/http"
	"path/filepath"
)

var files = make(map[string]string)

func main() {
	cfg, err := readConfig(filepath.Join(HOME, ".logcat"))
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range cfg.Logs {
		key := fmt.Sprintf("/%s", k)
		files[key] = v.Path
		http.HandleFunc(key, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, files[r.RequestURI])
		})
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}
