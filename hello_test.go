package hello_test

import (
	"bytes"
	"testing"

	"github.com/TsoiAlexx/hello"
)

func TestPrintPrintsHelloMessageToOutput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	p := hello.NewPrinter()
	p.Output = buf
	p.Print()
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
