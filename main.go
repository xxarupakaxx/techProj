package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w htto.ResponseWriter, r *http.Request) {
		//httpリクエストは論理パスなのでpathパッケージ
		if ok, err := path.Match("/data/*.txt", r.URL.Path); err != nil || !ok {
			http.NotFound(w, r)
			return
		}

		//以降はっパスを物理パスとして扱うのでpath/filepathパッケージ
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))
		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}
