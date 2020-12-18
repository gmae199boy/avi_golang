package handler

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"github.com/davecgh/go-spew/spew"
	"os"
	"os/exec"
)

type VideoHandler struct {}

func (h VideoHandler) UploadVideo(w http.ResponseWriter, r *http.Request) {
	f := new(FFMPEG)
	f.RunFfmpeg()
	fmt.Fprintf(w, "video convert complete")
}

type FFMPEG struct {

}

func (f * FFMPEG) RunFfmpeg() {
	cmd := "/usr/bin/env"
	args := []string{"--", "ffmpeg", "-i", "./assets/IMG_2245.MOV", "-y","i.avi"}
	c := exec.Command(cmd, args...)
	spew.Dump(c)
	c.Stderr = os.Stderr
    c.Stdout = os.Stdout
    // c.Stdin = os.Stdin
	c.Env = os.Environ()
	// c.SysProcAttr = &syscall.SysProcAttr{
	//   Pdeathsig: syscall.SIGTERM,
	// }
	_ = c.Run()
	fmt.Println("ffmpeg encoding is over!!!")
}