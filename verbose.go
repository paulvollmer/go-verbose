package verbose

import (
	"fmt"
	"io"
)

type Verbose struct {
	Active bool
	Writer io.Writer
}

func New(w io.Writer, a bool) Verbose {
	v := Verbose{}
	v.Active = a
	v.Writer = w
	return v
}

func (v *Verbose) Println(a ...interface{}) {
	if v.Active {
		fmt.Fprintln(v.Writer, a...)
	}
}

func (v *Verbose) Printf(format string, a ...interface{}) {
	if v.Active {
		fmt.Fprintf(v.Writer, format, a...)
	}
}
