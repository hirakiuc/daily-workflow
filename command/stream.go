package command

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Out io.Writer
}

type IoStream struct {
	Out *Printer
	Err *Printer
}

func NewPrinter(out io.Writer) *Printer {
	return &Printer{
		Out: out,
	}
}

func (p *Printer) Print(a ...interface{}) {
	fmt.Fprint(p.Out, a...)
}

func (p *Printer) Println(a ...interface{}) {
	fmt.Fprintln(p.Out, a...)
}

func (p *Printer) Printf(format string, a ...interface{}) {
	fmt.Fprintf(p.Out, format, a...)
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
