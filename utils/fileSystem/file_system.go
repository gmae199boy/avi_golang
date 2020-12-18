package fileSystem

import (
	"os"
	"io/ioutil"
	"log"
	"errors"
	// "github.com/davecgh/go-spew/spew"
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

func ReadNameOfFileList(path string) ([]string, error) {
    // f, err := os.Open(path)
    // if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	fullPath := "./assets/" + path
    files, err := ioutil.ReadDir(fullPath)
    // f.Close()
    if err != nil {
		log.Println(err)
		return nil, err
	}

	fileName := make([]string, len(files))
    for i, file := range files {
		// if file[0:1] == "." {
		// 	continue
		// }
		fileName[i] = file.Name()
	}
	log.Println(fileName)
	return fileName, nil
}