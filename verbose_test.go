package verbose

import (
	"bytes"
	"testing"
)

func Test_Verbose(t *testing.T) {
	var buf bytes.Buffer

	debug := New(&buf, true)
	if debug.Active != true {
		t.Error("Verbose New failed")
	}

	n, err := debug.Print("hello Print\n")
	if err != nil {
		t.Error("Verbose.Print failed")
	}
	if n != 12 {
		t.Error("Verbose.Print returned n not equal")
	}

	n, err = debug.Printf("hello Printf integer %v\n", 1337)
	if err != nil {
		t.Error("Verbose.Printf failed")
	}
	if n != 26 {
		t.Error("Verbose.Printf returned n not equal")
	}

	n, err = debug.Println("hello Println")
	if err != nil {
		t.Error("Verbose.Println failed")
	}
	if n != 14 {
		t.Error("Verbose.Println returned n not equal")
	}

	t.Log(&buf)
}
