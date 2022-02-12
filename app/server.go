package app

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"time"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 100 // 100MB

// The server only receives files
func RunServer(config *Config) error {
	fmt.Printf("Waiting for files at %s\n", config.ServerAddr)

	http.HandleFunc("/upload", handler)

	if err := http.ListenAndServe(config.ServerAddr, nil); err != nil {
		return err
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength > MAX_UPLOAD_SIZE {
		http.Error(w, "The uploaded image is too big. Please use an image less than 100MB in size", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	exts, err := mime.ExtensionsByType(r.Header.Get("Content-Type"))
	if err != nil {
		return
	}

	dst, err := os.Create(fmt.Sprintf("./%d%s", time.Now().UnixNano(), exts[len(exts)-1]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = dst.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received %s", dst.Name())
	fmt.Printf("Received %s\n", dst.Name())
}
