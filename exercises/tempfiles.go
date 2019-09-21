package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func main() {
	d := MakeTempDir()
	log.Print("created dir:", d)

	err := os.RemoveAll(d)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("removed dir:", d)
}

func MakeTempDir() (dirName string) {
	u, _ := user.Current()
	dirName, err := ioutil.TempDir(u.HomeDir, ".t")
	if err != nil {
		log.Fatal(err)
	}
	return
}
