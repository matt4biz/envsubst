package envsubst

import (
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	os.Setenv("BAR", "bar")
}

// Basic integration tests. because we  already test the
// templating processing in envsubst/parse;
func TestIntegration(t *testing.T) {
	input, expected := "foo $BAR", "foo bar"
	str, err := String(input)
	if str != expected || err != nil {
		t.Error("Expect string integration test to pass")
	}
	bytes, err := Bytes([]byte(input))
	if string(bytes) != expected || err != nil {
		t.Error("Expect bytes integration test to pass")
	}
	bytes, err = ReadFile("testdata/file.tmpl")
	fexpected, _ := ioutil.ReadFile("testdata/file.out")
	if string(bytes) != string(fexpected) || err != nil {
		t.Error("Expect ReadFile integration test to pass")
	}
	bytes, err = ReadFileSkipping("testdata/skipping.tmpl")
	sexpected, _ := ioutil.ReadFile("testdata/skipping.out")
	if string(bytes) != string(sexpected) || err != nil {
		t.Error("Expect ReadFileSkipping integration test to pass")
	}
}
