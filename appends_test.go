package appends

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAppendText(t *testing.T) {
	expected := []byte("cookies")

	err := AppendText("fixtures/blankText.txt", "cookies")

	actual, err := ioutil.ReadFile("fixtures/blankText.txt")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if bytes.Equal(expected, actual) == false {
		fmt.Errorf("The Appended File was append incorrectly")
		t.Fail()
	}
}
