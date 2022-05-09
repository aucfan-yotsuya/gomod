package scanner

import (
	"bufio"
	"io/ioutil"
	"os"
)

type Scanner struct {
	File    *os.File
	Scanner *bufio.Scanner
	Byte    []byte
}

var err error

func New() *Scanner {
	return new(Scanner)
}
func Once(filename string) []byte {
	var s *Scanner = New()
	defer s.File.Close()
	return s.Open(filename).ReadAll().Byte
}
func (s *Scanner) Open(filename string) *Scanner {
	if s.File, err = os.Open(filename); err != nil {
		panic(err)
	}
	return s
}
func (s *Scanner) notNilFile() bool {
	return s.File != nil
}
func (s *Scanner) Close() *Scanner {
	if s.notNilFile() {
		s.File.Close()
	}
	return s
}
func (s *Scanner) ReadAll() *Scanner {
	if s.Byte, err = ioutil.ReadAll(s.File); err != nil {
		panic(err)
	}
	return s
}
func (s *Scanner) NewScanner() *Scanner {
	if s.notNilFile() {
		s.Scanner = bufio.NewScanner(s.File)
	}
	return s
}
func (s *Scanner) notNilScanner() bool {
	return s.Scanner != nil
}
func (s *Scanner) Scan() bool {
	return s.notNilScanner() && s.Scanner.Scan()
}
func (s *Scanner) Text() string {
	var t = ""
	if s.notNilScanner() {
		t = s.Scanner.Text()
	}
	return t
}
func (s *Scanner) Bytes() []byte {
	var b = []byte("")
	if s.notNilScanner() {
		b = s.Scanner.Bytes()
	}
	return b
}
