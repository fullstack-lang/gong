// generated code - do not edit
package static

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"time"

	// this package contains ...
	"github.com/fullstack-lang/gong/lib/split"
)

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ServeStaticFiles(logFlag bool) (r *http.ServeMux) {
	r = http.NewServeMux()

	handler := EmbedFolder(split.NgDistNg, "ng-github.com-fullstack-lang-gong-lib-split/dist/ng-github.com-fullstack-lang-gong-lib-split/browser")
	r.Handle("/", handler)

	return r
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) http.Handler {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServer(http.FS(fsys))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		f, err := fsys.Open(path)
		if err != nil {
			indexFile, err := fsys.Open("index.html")
			if err != nil {
				http.NotFound(w, r)
				return
			}
			defer indexFile.Close()
			if seeker, ok := indexFile.(io.ReadSeeker); ok {
				http.ServeContent(w, r, "index.html", time.Time{}, seeker)
			} else {
				data, _ := io.ReadAll(indexFile)
				http.ServeContent(w, r, "index.html", time.Time{}, strings.NewReader(string(data)))
			}
			return
		}
		f.Close()

		fileServer.ServeHTTP(w, r)
	})
}

func RunServer(mux *http.ServeMux, addr string) error {
	return http.ListenAndServe(addr, CorsHandler(mux))
}
