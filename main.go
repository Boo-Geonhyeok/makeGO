package main

import (
	"flag"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	folderName := flag.String("fn", "", "folder name")

	flag.Parse()

	if *folderName == "" {
		log.Fatal("ENTER FOLDER NAME")
	}

	path := filepath.Join(currentUser.HomeDir, "go/src/github.com/Boo-Geonhyeok/", *folderName)
	filePath := path + "/main.go"

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	code := []byte("package main\n\nfunc main() {\n\n}")
	err = os.WriteFile(filePath, code, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
