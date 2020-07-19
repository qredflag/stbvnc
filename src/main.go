package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/key", HandleEventInjection)
	http.HandleFunc("/fb", HandleFBRequest)
	http.HandleFunc("/", HandleIndexRequest)

	_ = http.ListenAndServe(":8080", nil)
}

func HandleIndexRequest(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buffer := make([]byte, 65536)

	for {
		_, read_error := file.Read(buffer)
		if read_error != nil {
			if read_error != io.EOF {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			break
		}
		_, write_error := w.Write(buffer)
		if write_error != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

}

func HandleEventInjection(w http.ResponseWriter, r *http.Request) {
	param1, _ := r.URL.Query()["keystroke"]
	param2, _ := r.URL.Query()["type"]

	keystroke := strings.Join(param1, "")
	typep := strings.Join(param2, "")

	if len(keystroke) < 1 || (typep != "1" && typep != "0") {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	key := map[bool]string{true: "-ks", false: "-kqt"}[typep == "0"]

	cmdline := key + " " + keystroke

	fmt.Println(cmdline)

	cmd := exec.Command("/usr/local/share/app/bin/sendqtevent", key+" "+keystroke)
	cmd.Start()
}

func HandleFBRequest(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/dev/fb0")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("Content-Encoding", "gzip")
	w.Header().Add("Content-Disposition", "attachment; filename=\"fb.raw\"")

	buffer := make([]byte, 65536)

	writer, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer writer.Close()

	for {
		bytes_count, read_error := file.Read(buffer)
		if read_error != nil {
			if read_error != io.EOF {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			break
		}
		_, write_error := writer.Write(buffer[:bytes_count])
		if write_error != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

}
