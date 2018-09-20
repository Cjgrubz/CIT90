package main

import (
	"strings"
	"io"
	"log"
	"os"
	"fmt"
)

func main() {
	//create string
	//assign to var
	name := "James"
	
	str := `html here` + name + `more html`

	fmt.Println(str)

	//create string
	//string print

	s := fmt.Sprint(`maas `+ name + `menos`)

	nf, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}
	//Create file
	//io copy to the file

	io.Copy(nf, strings.NewReader(s))
	//Create File
	//write string to file

	nf2, err := os.Create("newfile2.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	n, err := nf2.WriteString(str)
	if err != nil {
		log.Fatal("whoops2", err)
	}
	
	fmt.Println("bytes written", n)
}