package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
)

var input_file string

func exist(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func parse() string {

	var (
		input    []byte
		out      *os.File
		err      error
		out_file string = "output.html"
	)
	// check out file exist
	if exist(input_file) {
		log.Fatal("File is not exist")
		os.Exit(-1)
	}
	// open markdown file
	if input, err = ioutil.ReadFile(input_file); err != nil {
		log.Fatal("Error of reading markdown file")
		os.Exit(-1)
	}

	// transform markdown to the html
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// create html file
	if out, err = os.Create(out_file); err != nil {
		log.Fatal("Error of creating html")
		os.Exit(-1)
	}

	defer func() {
		if err = out.Close(); err != nil {
			log.Fatal("Can't close the file")
			os.Exit(-1)
		}
	}()

	if _, err = out.Write(html); err != nil {
		log.Fatal("Error of writing html")
		os.Exit(-1)
	}

	return out_file
}

func command() (string, string) {
	fmt.Println("Usage")
	usage := "\tExecute : ./glue [inputfile] [port]\n\tHelp    : ./glue"

	if len(os.Args) <= 1 {
		fmt.Println(usage)
		os.Exit(-1)
	} else if len(os.Args) <= 2 && len(os.Args) > 1 {
		return string(os.Args[1]), "9090"
	}

	return string(os.Args[1]), string(os.Args[2])
}

func main() {
	var port string
	input_file, port = command()

	server(":" + port)
}
