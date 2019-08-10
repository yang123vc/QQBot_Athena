package models

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getText(file string) (text string) {
	path := "txt\\" + file
	text = ""
	menuFile, fileError := os.Open(path)
	if fileError != nil {
		text = "File Not Found"
		return
	}
	inputReader := bufio.NewReader(menuFile)
	for {
		inputString, inputError := inputReader.ReadString('\n')
		text += inputString
		if inputError == io.EOF {
			return
		}
	}
}

func ReadFile(fileName string) (fs []byte, err error) {
	/*
		fs, err = ioutil.ReadFile(fileName)
		return

	*/

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	return b, err
}

func WriteFilePre(fileName string) (writer *bufio.Writer, err error) {
	resPath := "resource\\"
	file, err := os.Create(resPath + fileName)
	if err != nil {
		return
	}

	writer = bufio.NewWriter(file)
	//err = ioutil.WriteFile(fileName, data, 0666)
	return
}

func GetWebFile(url string) (filePath string, err error) {
	res, err := http.Get(url)
	fmt.Println("res")
	if err != nil {
		return
	}
	defer res.Body.Close()

	reader := bufio.NewReaderSize(res.Body, 32*1024)

	fmt.Println("reader")
	//resPath := "\\resource\\"

	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	fileName := strconv.Itoa(rand1.Int()) + ".jpg"
	//fileName := path.Base(url)
	fmt.Println(fileName)
	writer, err := WriteFilePre(fileName)
	_, _ = io.Copy(writer, reader)
	//fmt.Print("length: %d", written)

	filePath = "resource\\" + fileName
	return
}
