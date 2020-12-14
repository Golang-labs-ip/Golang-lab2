package lab2

import (
	"io"
	"io/ioutil"
)

//ComputeHandler ...
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

//Compute ...
func (ch *ComputeHandler) Compute() error {
	expr, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	res := PostfixToPrefix(string(expr))

	buf := []byte(res)
	if _, e := ch.Output.Write(buf); e != nil {
		return e
	}

	return nil
}
