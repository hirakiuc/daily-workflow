package command

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIoStream(t *testing.T) {
	t.Parallel()

	s := NewIoStream()
	assert.NotNil(t, s)

	assert.NotNil(t, s.Out)
	assert.Equal(t, s.Out.out, os.Stdout)

	assert.NotNil(t, s.Err)
	assert.Equal(t, s.Err.out, os.Stderr)
}

func TestIoStreamSetOut(t *testing.T) {
	t.Parallel()

	s := NewIoStream()
	assert.NotNil(t, s)

	var buf bytes.Buffer

	assert.NotEqual(t, s.Out.out, &buf)

	s.SetOut(&buf)
	assert.Equal(t, s.Out.out, &buf)
}

func TestIoStreamSetErr(t *testing.T) {
	t.Parallel()

	s := NewIoStream()
	assert.NotNil(t, s)

	var buf bytes.Buffer

	assert.NotEqual(t, s.Err.out, &buf)

	s.SetErr(&buf)
	assert.Equal(t, s.Err.out, &buf)
}
