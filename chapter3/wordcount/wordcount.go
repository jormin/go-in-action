package main

import (
	"github.com/jormin/go-in-action/chapter3/words"
	"gitlab.wcxst.com/jormin/go-tools/log"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	log.Info("filename: %s", filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Read file [%s] error: %+v", filename, err)
	}
	count := words.CountWords(string(b))
	log.Info("File [%s] has [%d] word, content is: %s", filename, count, b)
}
