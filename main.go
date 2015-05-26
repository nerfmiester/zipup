package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	// First read the directory passed in as value

	if len(os.Args) < 1 {
		return
	}

	now := time.Now()

	flag := os.Args[2]

	writeFile := os.Args[1] + "/fileOut.txt"

	filew, err := os.OpenFile(writeFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer filew.Close()

	direcoryContents, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	for _, content := range direcoryContents {

		if content.IsDir() {
			fmt.Println("Is Directory ==> ", content.Name())
		} else {

			if content.ModTime().Day() == now.Add((24 * time.Hour * -6)).Day() {

				fmt.Println(content.Name())
				fmt.Println(content.ModTime())

				fullPath := os.Args[1] + "/" + content.Name()

				fileContent, err := ioutil.ReadFile(fullPath)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(fileContent))
				_, err = filew.WriteString(string(fileContent))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	switch flag {
	case "z":
		{
			err = zipit(writeFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	case "c":
		{
			err = compress(writeFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	case "f":
		{
			err = fizzbuzz(100)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

// Truth fred
type Truth func(i int) (fizz string) //

// Truthz fred
type Truthz func(i *int) // truth

// Trutha fred
type Trutha struct {
	counter int
}

// NewTrutha fred
func NewTrutha(counter int) *Trutha {
	a := new(Trutha)
	a.counter = counter
	return a
}
func (wf *Trutha) addOne(ix int) (x int) {
	return ix + 1
}
func (wf *Trutha) getTrutha() int {
	return wf.counter
}
func (wf *Trutha) getTruthaString() string {
	return strconv.Itoa(wf.counter)
}

// WriteFizz fred
//func (wf Truth) WriteFizz(i int) (fizz string) {
//	return wf(i)
//}
func myTruth(i int) (fizz string) {

	if i%15 == 0 {
		return "fizzbang"
	} else if i%3 == 0 {
		return "fizz"
	} else if i%5 == 0 {
		return "bang"
	}
	return strconv.Itoa(i)

}
func myTruthx(i int) (fizz string) {

	str := strconv.Itoa(i)

	if strings.ContainsAny(str, "7") {
		return "whoosh"
	}

	return str

}
func myTruthy(i int) (fizz string) {

	//Bytestr := []byte(strconv.Itoa(i))
	str := strconv.Itoa(i)

	r, _ := regexp.Compile("7")

	if r.MatchString(str) {
		return "whoosh"
	}

	return str

}
func myTruthz(i *int) {

	//Bytestr := []byte(strconv.Itoa(i))
	*i = *i + 10

}

func fizzbuzz(number int) error {

	for i := NewTrutha(1); i.getTrutha() < number; (*Trutha).addOne(i, i.counter) {

		//fmt.Println(Truth(myTruthx)(i))
		//fmt.Println(Truth(myTruth)(i))
		//fmt.Println(Truth(myTruthy)(i))

		//Truthz(myTruthz)(&i)
		fmt.Printf("Welcome -> %s\n", i.getTruthaString())

	}

	return nil

}

func zipit(writeFile string) error {

	newfile, err := os.Create(writeFile + ".zip")
	if err != nil {
		fmt.Println(err)
	}
	defer newfile.Close()
	zipit := zip.NewWriter(newfile)
	defer zipit.Close()

	zipfile, err := os.Open(writeFile)
	if err != nil {
		fmt.Println(err)
	}
	defer zipfile.Close()

	// get the file information
	info, err := zipfile.Stat()
	fmt.Println("Zip file info ==> " + info.Name())
	if err != nil {
		fmt.Println(err)
	}

	header, err := zip.FileInfoHeader(info)
	fmt.Println("Zip file header ==> " + header.ModTime().String())
	if err != nil {
		fmt.Println(err)
	}

	writer, err := zipit.CreateHeader(header)

	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(writer, zipfile)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	return nil

}

func compress(writeFile string) error {

	fmt.Println("writeFile ==> " + writeFile)

	//newfile, err := os.Create(writeFile + ".gz")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer newfile.Close()

	rawfile, err := os.Open(writeFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// calculate the buffer size for rawfile
	info, _ := rawfile.Stat()

	fmt.Println(reflect.TypeOf(info))

	var size int64 = info.Size()
	rawbytes := make([]byte, size)

	// read rawfile content into buffer
	buffer := bufio.NewReader(rawfile)
	_, err = buffer.Read(rawbytes)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	writer.Write(rawbytes)
	writer.Close()

	err = ioutil.WriteFile(writeFile+".gz", buf.Bytes(), info.Mode())
	// use 0666 to replace info.Mode() if you prefer

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileComp, err := os.OpenFile(writeFile+".gz", os.O_RDONLY, 0660)
	if err != nil {
		fmt.Println(err)
	}
	defer fileComp.Close()

	fileCompInfo, _ := fileComp.Stat()

	fmt.Printf("%s compressed to %s original size was %d now shrunk to %d \n", writeFile, writeFile+".gz", info.Size(), fileCompInfo.Size())

	return nil

}
