package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"

	lab2 "github.com/Golang-labs-ip/Golang-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	writeFile       = flag.String("o", "", "Write result to file")
	readFile        = flag.String("f", "", "Read expression from file")
)

func main() {

	flag.Parse()

	var (
		from io.Reader
		to   io.Writer
		err  error
	)

	switch {
	case *inputExpression != "":
		from = strings.NewReader(*inputExpression)
		break
	case *readFile != "":
		data, err := ioutil.ReadFile(*readFile)
		if err != nil {
			os.Stderr.WriteString("cant find the file")
			return
		}
		from = strings.NewReader(string(data))
		break
	default:
		os.Stderr.WriteString("can't find the expression")
		return
	}

	if *writeFile != "" {
		if to, err = os.Create(*writeFile); err != nil {
			os.Stderr.WriteString("error with creating a file")
		}
	} else {
		to = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  from,
		Output: to,
	}
	if err := handler.Compute(); err != nil {
		os.Stderr.WriteString(err.Error())
	}

}
