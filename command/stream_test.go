package command_test

import (
	"bytes"
	"fmt"
	"testing"

	base "github.com/hirakiuc/daily-workflow/command"
	"github.com/stretchr/testify/assert"
)

func TestNewPrinter(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	p := base.NewPrinter(&buf)
	assert.NotNil(t, p)
}

func TestPrinterPrint(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	p := base.NewPrinter(&buf)
	assert.NotNil(t, p)

	const msg = "print test\n"

	p.Print(msg)

	str := buf.String()
	assert.Equal(t, str, msg)
}

func TestPrinterPrintln(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	p := base.NewPrinter(&buf)
	assert.NotNil(t, p)

	const msg = "print test without new line"

	p.Println(msg)

	str := buf.String()
	assert.Equal(t, str, msg+"\n")
}

func TestPrinterPrintf(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	p := base.NewPrinter(&buf)
	assert.NotNil(t, p)

	const format = "printf test:val(%s)\n"

	const val = "sample value"

	p.Printf(format, val)

	str := buf.String()
	expected := fmt.Sprintf(format, val)
	assert.Equal(t, str, expected)
}
