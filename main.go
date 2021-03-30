package main

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"unicode"
	"strings"
)

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func binaryToString(s string) string {
	binarys := strings.Fields(s)
	out := make([]byte, len(binarys))
	for i := 0; i < len(binarys); i ++ {
		b, err := strconv.ParseInt(string(binarys[i]), 2, 64)
		if err != nil {
			panic(err)
		}
		out[i] = byte(b)
	}
	return string(out)
}

func writeFile(content string, name string){
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File created ", name)
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	text := ""
	binary := ""
	read, err := zip.OpenReader(os.Args[1])
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for i:=0; i<len(read.File); i++ {
		str := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(i)))
		opena, _ := read.Open(str)
		content, _ := ioutil.ReadAll(opena)
		binary = binary +" "+ string(content)
		text = text + binaryToString(string(content))
	}
	writeFile(text, "kartaca.txt")
	writeFile(binary, "kartaca-binary.txt")
}