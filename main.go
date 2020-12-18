package main

import (
    "context"
	// "fmt"
	"log"
	// "time"
    "net/http"
    "runtime"
    "github.com/gmae199boy/avi_golang/router"
    "github.com/joho/godotenv"
)

// const maxUploadSize = 2 * 1024 * 1024 // 2 mb
// const uploadPath = "./assets"

func main() {
    godotenv.Load()
    runtime.GOMAXPROCS(4)
    context.Background()

	// fs := http.FileServer(http.Dir(uploadPath))
    // http.Handle("/video/", http.StripPrefix("/video", fs))

    // // All handling url : video
    // http.HandleFunc("/v1/video", router.VideoRouter) 

    // // All handling url : file
    http.HandleFunc("/v1/folder", router.FolderRouter)

    // // All handling url : user
    // http.HandleFunc("/v1/user", router.UserRouter)

	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
    log.Fatal(http.ListenAndServe(":8080", nil))
}