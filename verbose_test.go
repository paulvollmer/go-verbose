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
	debug.Println("hello log")
	debug.Printf("hello integer %v\n", 1337)

	t.Log(&buf)
}
