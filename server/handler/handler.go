package handler

import (
	"net/http"
	"strings"
	"sync/atomic"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// GetIndex it will display index.html when it requested
func GetIndex(directory string) http.Handler {
	return http.FileServer(FileSystem{http.Dir(directory)})
}

// GetStatic it will display the statics when it requested
// Im using a custom http.FileSystem to avoid display the whole folder
func GetStatic(directory string) http.Handler {
	return http.StripPrefix("/static/", GetIndex(directory))
}

// GetHealt check the healt of the server
func GetHealt(healthy *int32) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
