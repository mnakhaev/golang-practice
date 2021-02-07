package main

import (
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
3
4
5
`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {

	in := strings.NewReader(testOk)
	out := new(bytes.Buffer) // just new buffer
	if err := uniq(in, out); err != nil {
		t.Errorf("test for OK was failed")
	}

	result := out.String()
	if result != testOkResult {
		t.Errorf("test for OK failed - results don't match\n %v %v", result, testOkResult)
	}
}

var testFail = `1
4
2
3
5
2
`

func TestForError(t *testing.T) {
	in := strings.NewReader(testFail)
	out := new(bytes.Buffer) // just new buffer
	if err := uniq(in, out); err == nil {
		t.Errorf("test for Fail was passed")
	}

}
