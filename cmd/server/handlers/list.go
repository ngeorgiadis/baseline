package handlers

import (
	"baseline/m/v2/cmd/server/config"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"path"
	"path/filepath"

	"github.com/gorilla/mux"
)

// DirectoryContent ...
type DirectoryContent struct {
	FullPath string         `json:"full_path"`
	Files    []PlaylistFile `json:"files"`
}

// PlaylistFile ...
type PlaylistFile struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
	Type string `json:"type"`
}

func md5FromString(s string) string {
	k := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", k)
}

func listHandler(c *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		q := r.URL.Query()
		basePath := filepath.ToSlash(c.App.RootPath)
		if _, ok := q["p"]; ok {
			basePath = path.Join(basePath, q.Get("p"))
		}

		list, err := ioutil.ReadDir(basePath)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		files := []PlaylistFile{}
		for _, f := range list {
			t := "FILE"
			if f.IsDir() {
				t = "DIR"
			}
			files = append(files, PlaylistFile{
				Name: f.Name(),
				Hash: md5FromString(f.Name()),
				Type: t,
			})
		}
		b, _ := json.Marshal(files)
		w.Write(b)
	}
}

func fileHandler(c *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		q := r.URL.Query()
		basePath := filepath.ToSlash(c.App.RootPath)
		if _, ok := q["p"]; ok {
			basePath = path.Join(basePath, q.Get("p"))
		}

		fPath := path.Join(basePath, vars["filename"])
		b, _ := ioutil.ReadFile(fPath)
		w.Write(b)
	}
}
