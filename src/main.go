package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/key", HandleEventInjection)
	http.HandleFunc("/fb", HandleFBRequest)
	_ = http.ListenAndServe(":8080", nil)
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

const SCREEN_WIDTH = 1920
const SCREEN_HEIGHT = 1080

func HandleFBRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "image/png")
	w.Header().Add("Content-Disposition", "attachment; filename=\"fb.png\"")

	file, err := os.Open("/dev/fb0")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	img := image.NewNRGBA(image.Rect(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT))
	io.LimitReader(file, SCREEN_WIDTH*SCREEN_HEIGHT*4).Read(img.Pix)

	//convert BGRA to RGBA
	for offset := 0; offset < SCREEN_WIDTH*SCREEN_HEIGHT*4; offset += 4 {
		img.Pix[offset], img.Pix[offset+2] = img.Pix[offset+2], img.Pix[offset]
	}

	enc := &png.Encoder{CompressionLevel: png.BestSpeed}

	encode_error := enc.Encode(w, img)
	if encode_error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
