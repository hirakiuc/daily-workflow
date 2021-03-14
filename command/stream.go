package command

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	out io.Writer
}

type IoStream struct {
	Out *Printer
	Err *Printer
}

func NewPrinter(out io.Writer) *Printer {
	return &Printer{
		out: out,
	}
}

func (p *Printer) Print(a ...interface{}) {
	fmt.Fprint(p.out, a...)
}

func (p *Printer) Println(a ...interface{}) {
	fmt.Fprintln(p.out, a...)
}

func (p *Printer) Printf(format string, a ...interface{}) {
	fmt.Fprintf(p.out, format, a...)
}

func NewIoStream() *IoStream {
	return &IoStream{
		Out: NewPrinter(os.Stdout),
		Err: NewPrinter(os.Stderr),
	}
}

func (s *IoStream) SetOut(v io.Writer) {
	s.Out = NewPrinter(v)
}

func (s *IoStream) SetErr(v io.Writer) {
	s.Err = NewPrinter(v)
}
