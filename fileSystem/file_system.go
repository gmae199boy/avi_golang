package fileSystem

import (
	"os"
	"log"
	"errors"
)

func CreateFolder(path string, mode os.FileMode) error {
	fullPath := "./assets/" + path
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.Mkdir(fullPath, mode)
		log.Println("success create folder to " + fullPath)
		return nil
	}
	return errors.New("fileSystem Error: File already exist")
}

func DeleteFolder(path string) error {
	fullPath := "./assets/" + path
	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		err = os.RemoveAll(fullPath)
		if err != nil {
			log.Println("FileSystem Error: Folder delete was failed")
			return errors.New("FileSystem Error: Folder delete was failed")
		}
		log.Println("FileSystem: file was deleted!", fullPath)
		return nil
	}
	return errors.New("FileSystem Error: prociser has gone")
}

func OpenFolder(path string) error {
    f, err := os.Open(path)
    if err != nil {
        log.Println(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Println(err)
    }

    for _, file := range files {
        log.Println(file.Name())
	}
	return errors.New("asd")
}