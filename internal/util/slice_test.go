package util

import (
	"os"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesInt(t *testing.T) {
	duplicate := []int32{1, 1, 1, 2, 2, 2, 3, 3}
	result := []int32{1, 2, 3}
	unique := RemoveDuplicatesInt(duplicate)
	if !reflect.DeepEqual(unique, result) {
		t.Errorf("RemoveDuplicatesInt: expected %v but got %v", result, unique)
	}
}

func TestStrToint(t *testing.T) {
	stringInts := []string{"1", "2", "3", "4"}
	result := []int{1, 2, 3, 4}
	convertedSlice, _ := StrToint[int](stringInts)

	if !reflect.DeepEqual(convertedSlice, result) {
		t.Errorf("StrToint: expected %v but got %v", result, convertedSlice)
	}
}

func TestStrTointFailingCase(t *testing.T) {
	// Test case: Input contains a non-numeric string
	input := []string{"42", "invalid", "100"}

	_, err := StrToint[int](input)

	if err == nil {
		t.Errorf("StrToint: expected error but got nil")
	}
}

func TestReadStdin(t *testing.T) {
	fakeInput := []byte("1\n2\n3\n")
	expected := []string{"1", "2", "3"}
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write(fakeInput); err != nil {
		t.Error(err)
	}
	w.Close()

	// retore stdin right after the test
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	result := *ReadStdin()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
