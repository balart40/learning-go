package chapterone

import (
	"fmt"
	"testing"
)

func Test_chap_1_ex_1(t *testing.T) {
	result := helloWorld()
	if result != "Hello World" {
		t.Error("incorrect result: expected Hello World, got", result)
	} else {
		fmt.Println("Hello world Passed! Success!")
	}
}
