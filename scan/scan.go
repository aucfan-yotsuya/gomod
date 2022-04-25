package scan

import (
	"bufio"
	"io/ioutil"
	"os"
)

type Scan struct {
	File    *os.File
	Scanner *bufio.Scanner
	Byte    []byte
}

var err error

func New() *Scan {
	return new(Scan)
}
func Once(filename string) []byte {
	var s *Scan = New()
	defer s.File.Close()
	return s.Open(filename).ReadAll().Byte
}
func (s *Scan) Open(filename string) *Scan {
	if s.File, err = os.Open(filename); err != nil {
		panic(err)
	}
	return s
}
func (s *Scan) Close() *Scan {
	s.File.Close()
	return s
}
func (s *Scan) ReadAll() *Scan {
	if s.Byte, err = ioutil.ReadAll(s.File); err != nil {
		panic(err)
	}
	return s
}
func (s *Scan) NewScanner() *Scan {
	s.Scanner = bufio.NewScanner(s.File)
	return s
}
